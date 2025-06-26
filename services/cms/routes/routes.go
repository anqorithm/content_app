package routes

import (
	"github.com/amal-meer/content_app/services/cms/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/upload-url", handlers.RequestUploadURL)
	app.Patch("/content/:id/status", handlers.UpdateContentStatus)
}
