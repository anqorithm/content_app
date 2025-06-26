package handlers

import (
	"context"
	"os"
	"time"

	"github.com/amal-meer/content_app/database"
	"github.com/amal-meer/content_app/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func GetAllContent(c *fiber.Ctx) error {
	var contents []models.Content
	if err := database.DB.Db.Order("publication_date desc").Find(&contents).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(contents)
}

func GetContentURL(c *fiber.Ctx) error {
	id := c.Params("id")

	var content struct {
		S3Key string
	}
	if err := database.DB.Db.Table("contents").Select("s3_key").Where("id = ?", id).First(&content).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Content not found")
	}

	url, err := GeneratePresignedDownloadURL(content.S3Key)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"download_url": url})
}

func GeneratePresignedDownloadURL(key string) (string, error) {
	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_S3_BUCKET")

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	client := s3.NewFromConfig(cfg)
	presign := s3.NewPresignClient(client)

	req, err := presign.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))
	return req.URL, err
}
