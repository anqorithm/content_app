package main

import (
	"fmt"

	"github.com/amal-meer/content_app/config"
	"github.com/amal-meer/content_app/database"
	"github.com/amal-meer/content_app/services/discovery/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize configuration
	config.InitConfig()
	
	// Initialize database
	database.ConnectDb()
	
	// Initialize a new Fiber app
	app := fiber.New()
	routes.SetupRoutes(app)
	
	// Use port 3001 for discovery service
	serverAddr := fmt.Sprintf("%s:3001", config.AppConfig.Server.Host)
	app.Listen(serverAddr)
}
