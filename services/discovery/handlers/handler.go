package handlers

import (
	"github.com/amal-meer/content_app/database"
	"github.com/amal-meer/content_app/models"
	"github.com/amal-meer/content_app/utils"
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

	url, err := utils.GeneratePresignedDownloadURL(content.S3Key)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"download_url": url})
}

