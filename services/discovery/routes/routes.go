package routes

import (
	"github.com/amal-meer/content_app/services/discovery/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/", "./frontend")
	app.Get("/contents", handlers.GetAllContent)
	app.Get("/content/:id/url", handlers.GetContentURL)
}
