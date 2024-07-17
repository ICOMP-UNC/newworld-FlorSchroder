package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "github.com/ICOMP-UNC/newworld-FlorSchroder/docs"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/database"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/routes"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/services"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:3000
// @BasePath /
// @termsOfService http://swagger.io/terms/
// @contact.email you@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	// Connect to the database
	dbPool, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	// Initialize the database
	err = database.InitDB(dbPool)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Inject the dbPool into your controllers as needed
	services.SetDB(dbPool)

	// Initialize the routes
	routes.InitRoutes(app, dbPool)

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}
