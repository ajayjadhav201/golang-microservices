package database

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// github.com/aws/aws-sdk-go/service/dynamodb
// github.com/aws/aws-sdk-go/aws/session

func GetConnection() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	return dynamodb.New(sess)
}

type Database interface {
}

type database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

func NewDatabase(connection *dynamodb.DynamoDB, logMode bool) Database {
	return &database{connection, logMode}
}

func (db *database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (d *database) FindAll() {
	//
}

func (db *database) FindOne(TableName string, condition map[string]interface{}) (*dynamodb.GetItemOutput, error) {

	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key:       conditionParsed,
	}

	return db.connection.GetItem(input)

}

func (db *database) CreateOrUpdate(TableName string, entity interface{}) (*dynamodb.PutItemOutput, error) {
	//
	item, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      item,
	}

	return db.connection.PutItem(input)
}

func (db *database) Delete(TableName string, condition map[string]interface{}) (*dynamodb.DeleteItemOutput, error) {
	//
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	//
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(TableName),
		Key:       conditionParsed,
	}
	//
	db.connection.DeleteItem(input)
	return nil, nil
}
