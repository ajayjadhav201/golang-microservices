package database

import (
	"context"
	"errors"
	"user-service/model"

	"golang-microservices/common"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDb struct {
	Client *dynamodb.Client
	health bool
}

func NewDynamoDb() UserStore {
	region := common.EnvString("AWS_REGION", "")
	accessKey := common.EnvString("AWS_ACCESS_KEY", "")
	secretKey := common.EnvString("AWS_SECRET_KEY", "")
	bucketName := common.EnvString("AWS_BUCKET_NAME", "")

	if region == "" || accessKey == "" || secretKey == "" || bucketName == "" {
		panic("one or more Environmental Variables not available")
	}
	db := dynamodb.New(dynamodb.Options{
		Region: region,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				accessKey, secretKey, "",
			),
		),
	})
	var tables *dynamodb.ListTablesInput
	_, err := db.ListTables(context.TODO(), tables)
	health := true
	if err != nil {
		health = false
		common.Println("database health is not good", err.Error())
	} else {
		createTable(db)
		common.Println("ajaj table created ")
	}
	return &DynamoDb{
		Client: db,
		health: health,
	}
}

func createTable(client *dynamodb.Client) {
	// Create table users
	tableName := "users"
	result, err := client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
	if err == nil {
		common.Println("ajaj table is already created: ", *result.Table.TableName)
		return
	}

	if !isResourceNotFoundException(err) {
		// If it's a different error, handle it
		common.Fatalf("failed to describe table: %v", err)
		return
	}

	common.Println("ajaj creating table: ", tableName)

	param := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeN,
			},
			{
				AttributeName: aws.String("Email"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("MobileNumber"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("EmailIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("Email"),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
			{
				IndexName: aws.String("MobileNumberIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("MobileNumber"),
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},
	}

	// Table doesn't exist, so create it
	_, err = client.CreateTable(context.TODO(), param)
	if err != nil {
		common.Fatalf("unable to create table error: %s", err.Error())
	}
}

func (db *DynamoDb) Health() bool {
	_, err := db.Client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	return err == nil
}

// isResourceNotFoundException checks if the error is a ResourceNotFoundException
func isResourceNotFoundException(err error) bool {
	var rnfe *types.ResourceNotFoundException
	if ok := errors.As(err, &rnfe); ok {
		return true
	}
	return false
}

func (db *DynamoDb) isUserExists(email string, mobileNumber string) (int, error) {
	count := 0

	// check email is registered
	if email != "" {
		emailQueryInput := &dynamodb.QueryInput{
			TableName:              aws.String("users"),
			IndexName:              aws.String("EmailIndex"),
			KeyConditionExpression: aws.String("Email = :email"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":email": &types.AttributeValueMemberS{Value: email},
			},
		}
		emailResult, err := db.Client.Query(context.TODO(), emailQueryInput)
		if err != nil {
			return count, err
		}

		if len(emailResult.Items) > 0 {
			count = count + 1
		}

	}

	// check mobile number is registered
	if mobileNumber != "" {
		mobileQueryInput := &dynamodb.ScanInput{
			TableName:        aws.String("users"),
			FilterExpression: aws.String("MobileNumber = :mobileNumber"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":mobileNumber": &types.AttributeValueMemberS{Value: mobileNumber},
			},
		}
		mobileResult, err := db.Client.Scan(context.TODO(), mobileQueryInput)
		if err != nil {
			return count, err
		}

		if len(mobileResult.Items) > 0 {
			count = count + 2
		}
	}

	return count, nil
}

// ConditionalCheckFailedException
// ProvisionedThroughputExceededException

func (db *DynamoDb) FindOne(TableName string, condition map[string]interface{}) (*dynamodb.GetItemOutput, error) {
	//
	conditionParsed, err := attributevalue.MarshalMap(condition)
	if err != nil {
		return nil, err
	}

	//
	input := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key:       conditionParsed,
	}
	return db.Client.GetItem(context.TODO(), input)
}

func (db *DynamoDb) Create(TableName string, entity interface{}, email string, mobile string) (*dynamodb.PutItemOutput, error) {

	// check if user is already exist
	ok, err := db.isUserExists(email, mobile)
	if err != nil {
		return nil, err
	} else if ok == 1 {
		return nil, &common.EmailAlreadyRegistered{Email: email}
	} else if ok == 2 {
		return nil, &common.MobileNumberAlreadyRegistered{MobileNumber: mobile}
	} else if ok == 3 {
		return nil, &common.EmailMobileAlreadyRegistered{Email: email, MobileNumber: mobile}
	}

	//insert user into database
	entityParsed, err := attributevalue.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(TableName),
	}
	return db.Client.PutItem(context.TODO(), input)
}

func (db *DynamoDb) Update(TableName string, entity interface{}) (*dynamodb.PutItemOutput, error) {
	//
	entityParsed, err := attributevalue.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	common.Println("ajaj entity is: ", entity)
	mail := entityParsed["Email"]
	common.Println("ajaj email value is :", mail)
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(TableName),
	}
	return db.Client.PutItem(context.TODO(), input)
}

func (db *DynamoDb) FindAll(TableName string) (*dynamodb.ScanOutput, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
		// Limit:     aws.Int32(10),
	}
	return db.Client.Scan(context.TODO(), input)
}

func (db *DynamoDb) Delete(TableName string, condition map[string]interface{}) (*dynamodb.DeleteItemOutput, error) {
	//
	conditionParsed, err := attributevalue.MarshalMap(condition)
	if err != nil {
		return nil, err
	}

	//
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(TableName),
		Key:       conditionParsed,
	}
	return db.Client.DeleteItem(context.TODO(), input)
}

func (db *DynamoDb) QueryEmail(TableName string, email string) (*dynamodb.QueryOutput, error) {
	//
	QueryInput := &dynamodb.QueryInput{
		TableName:              aws.String("users"),
		IndexName:              aws.String("EmailIndex"),
		KeyConditionExpression: aws.String("Email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
	}
	return db.Client.Query(context.TODO(), QueryInput)
}
func (db *DynamoDb) QueryMobileNumber(TableName string, mobile string) (*dynamodb.QueryOutput, error) {
	//
	QueryInput := &dynamodb.QueryInput{
		TableName:              aws.String("users"),
		IndexName:              aws.String("MobileNumberIndex"),
		KeyConditionExpression: aws.String("MobileNumber = :mobile"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":mobile": &types.AttributeValueMemberS{Value: mobile},
		},
		AttributesToGet: []string{
			"Password",
		},
		Limit: aws.Int32(1),
	}
	return db.Client.Query(context.TODO(), QueryInput)
}

//
//
//
//
//
//
//
//
//
//
//
// fucntions for handling user service

func (db *DynamoDb) CreateUser(user *model.User) (string, error) {
	common.Println("user going to get stored in database: ", user.ID)
	_, err := db.Create("users", *user, user.Email, user.MobileNumber)
	if err != nil {
		return "", err
	}
	common.Println("ajaj user inserted into database: ", user.ID)
	return common.Int64toa(user.ID), nil
}

func (db *DynamoDb) GetUsers() []*model.User {
	return []*model.User{}
}

func (db *DynamoDb) GetUserById(id string) (*model.User, error) {
	//id in database is int so convert string to int
	intid := common.Atoi(id)
	if intid == -1 {
		return nil, common.Error("Invalid id")
	}

	response, err := db.FindOne("users", map[string]interface{}{"ID": intid})

	if isResourceNotFoundException(err) {
		return nil, common.Error("User is not registered.")
	}
	if err != nil {
		return nil, err
	}
	//
	user := &model.User{}
	err = attributevalue.UnmarshalMap(response.Item, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DynamoDb) GetUserByEmailorMobile(emailorMobile string) (*model.User, error) {

	var response *dynamodb.QueryOutput
	var err error

	if common.IsEmail(emailorMobile) {
		response, err = db.QueryEmail("users", emailorMobile)
	} else if common.IsMobileNumber(emailorMobile) {
		response, err = db.QueryMobileNumber("users", emailorMobile)
	} else {
		return nil, common.Error("Please enter a valid mobile or email")
	}
	if isResourceNotFoundException(err) {
		return nil, common.Error("User is not registered.")
	}
	if err != nil {
		return nil, err
	}
	if len(response.Items) == 0 {
		return nil, common.Error("User is not registered.")
	}
	//
	user := &model.User{}
	err = attributevalue.UnmarshalMap(response.Items[0], user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DynamoDb) UpdateUser(id string, user *model.User) (*model.User, error) {
	// get previous user data
	common.Println("ajaj updateuserdtat is id: ", id, " and user is: ", user)
	u, err := db.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if isResourceNotFoundException(err) {
		return nil, common.Error("User is not registered.")
	}
	// merge data from latest to previous struct
	created := u.CreatedAt
	common.Println("ajaj in db previous user is: ", u)
	common.MergeStructs(u, user)
	u.CreatedAt = created
	//
	response, err := db.Update("users", u)
	if err != nil {
		return nil, err
	}
	// if len(response.Attributes) == 0 {
	// 	return nil, common.Error("User is not registered.")
	// }
	updatedUser := &model.User{}
	err = attributevalue.UnmarshalMap(response.Attributes, updatedUser)
	if err != nil {
		common.Println("ajaj error while unmarshaling json: ", err.Error())
		return nil, err
	}

	return updatedUser, nil
}

func (db *DynamoDb) DeleteUser(id string) error {
	return nil
}

/*
 // Item to insert
    item := map[string]types.AttributeValue{
        "id":             &types.AttributeValueMemberN{Value: "1"}, // id as uint64, converted to string
        "FirstName":      &types.AttributeValueMemberS{Value: "John"},
        "LastName":       &types.AttributeValueMemberS{Value: "Doe"},
        "Email":          &types.AttributeValueMemberS{Value: "john.doe@example.com"},
        "Password":       &types.AttributeValueMemberS{Value: "supersecret"},
        "MobileNumber":   &types.AttributeValueMemberS{Value: "1234567890"},
        "Address":        &types.AttributeValueMemberS{Value: "123 Elm St"},
        "ShippingAddress": &types.AttributeValueMemberS{Value: "456 Oak St"},
        "Token":          &types.AttributeValueMemberS{Value: "some-token"},
        "ProfileImage":   &types.AttributeValueMemberS{Value: "profile-image-url"},
        "CreatedAt":      &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
        "UpdatedAt":      &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
    }
*/
