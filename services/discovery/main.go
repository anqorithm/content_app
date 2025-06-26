package main

import (
	"github.com/amal-meer/content_app/database"
	"github.com/amal-meer/content_app/services/discovery/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	database.ConnectDb()
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3001")
}
