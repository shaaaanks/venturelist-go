package main

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	s3Region = "eu-west-2"
	s3Bucket = "venturelist-dev-bucket"
)

func uploadFile(filename string, file multipart.File) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3Region),
	})
	if err != nil {
		return "", fmt.Errorf("Error creating session: %v", err)
	}

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("Unable to upload %q to %q, %v", filename, s3Bucket, err)
	}

	return result.Location, nil
}
