package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/amal-meer/content_app/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GeneratePresignedUploadURL(filename string) (string, string, error) {
	storageConfig := config.AppConfig.Storage

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(storageConfig.AccessKey, storageConfig.SecretKey, "")),
		awsconfig.WithRegion(storageConfig.Region),
	)
	if err != nil {
		return "", "", err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(storageConfig.Endpoint)
	})
	presignClient := s3.NewPresignClient(client)

	key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), filename)
	req, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(storageConfig.Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(60*time.Minute))
	
	if err != nil {
		return "", "", err
	}
	
	return req.URL, key, err
}

func GeneratePresignedDownloadURL(key string) (string, error) {
	storageConfig := config.AppConfig.Storage

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(storageConfig.AccessKey, storageConfig.SecretKey, "")),
		awsconfig.WithRegion(storageConfig.Region),
	)
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(storageConfig.Endpoint)
	})
	presign := s3.NewPresignClient(client)

	req, err := presign.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(storageConfig.Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(60*time.Minute))
	
	if err != nil {
		return "", err
	}
	
	return req.URL, err
}