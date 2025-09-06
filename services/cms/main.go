package main

import (
	"fmt"

	"github.com/amal-meer/content_app/config"
	"github.com/amal-meer/content_app/database"
	"github.com/amal-meer/content_app/services/cms/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize configuration
	config.InitConfig()
	
	// Initialize database
	database.ConnectDb()
	
	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
	}))

	routes.SetupRoutes(app)
	
	serverAddr := fmt.Sprintf("%s:%s", config.AppConfig.Server.Host, config.AppConfig.Server.Port)
	app.Listen(serverAddr)
}
