package main

import (
	"github.com/amal-meer/content_app/database"
	"github.com/amal-meer/content_app/services/cms/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize a new Fiber app
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
		AllowMethods: "*",
	}))

	routes.SetupRoutes(app)
	app.Listen(":3000")
}
