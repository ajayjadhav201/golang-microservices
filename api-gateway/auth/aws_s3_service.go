package auth

import (
	"context"
	"mime/multipart"

	"github.com/ajayjadhav201/common"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsS3Service struct {
	BucketName string
	client     *s3.Client
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

func (awsService *AwsS3Service) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	//
	_, err = awsService.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(awsService.BucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   file, //here upload multipart file
		// ContentLength: aws.Int64(size),
		// ContentType: aws.String("text/plain"),        //text/plain  //image/jpeg
		// ACL: types.ObjectCannedACLPublicRead, // Make the object publicly readable
	})
	//
	// url := common.Sprintf("https://%s.s3.amazonaws.com/%s", BucketName, fileHeader.Filename)
	return fileHeader.Filename, err
}

func (awsService *AwsS3Service) DeleteFile(FileName string) error {
	//
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(awsService.BucketName),
		Key:    aws.String(FileName),
	}

	_, err := awsService.client.DeleteObject(context.Background(), input)
	return err
}
