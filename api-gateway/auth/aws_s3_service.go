package auth

import (
	"context"
	"mime/multipart"

	"golang-microservices/common"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsS3Service struct {
	BucketName string
	Uploader   *manager.Uploader
}

func NewAwsS3Service() *AwsS3Service {
	region := common.EnvString("AWS_REGION", "")
	accessKey := common.EnvString("AWS_ACCESS_KEY", "")
	secretKey := common.EnvString("AWS_SECRET_KEY", "")
	bucketName := common.EnvString("AWS_BUCKET_NAME", "")

	if region == "" || accessKey == "" || secretKey == "" || bucketName == "" {
		panic("one or more Environmental Variables not available")
	}
	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }
	// Create an Amazon S3 service client
	client := s3.New(
		s3.Options{
			Region: region,
			Credentials: aws.NewCredentialsCache(
				credentials.NewStaticCredentialsProvider(
					accessKey, secretKey, "",
				),
			),
		},
	)
	uploader := manager.NewUploader(client)
	//
	return &AwsS3Service{
		BucketName: bucketName,
		Uploader:   uploader,
	}
}

func UploadFile(Uploader *manager.Uploader, BucketName string, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	//
	_, err = Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   file, //here upload multipart file
		// ContentLength: aws.Int64(size),
		// ContentType: aws.String("text/plain"),        //text/plain  //image/jpeg
		// ACL: types.ObjectCannedACLPublicRead, // Make the object publicly readable
	})

	if err != nil {
		return "", err
	}
	// url := common.Sprintf("https://%s.s3.amazonaws.com/%s", BucketName, fileHeader.Filename)
	return fileHeader.Filename, nil
}
