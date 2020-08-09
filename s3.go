package main

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploadFile(filename string, file multipart.File) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.AwsRegion),
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyID, config.AwsSecretAccessKey, ""),
	})
	if err != nil {
		return "", fmt.Errorf("Error creating session: %v", err)
	}

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.AwsS3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("Unable to upload %q to %q, %v", filename, config.AwsS3Bucket, err)
	}

	return result.Location, nil
}
