package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "github.com/ICOMP-UNC/newworld-FlorSchroder/docs"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/analytics"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/controllers"
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
	dbPool, err := analytics.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	// Initialize the database
	err = analytics.InitDB(dbPool)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Inject the dbPool into your controllers as needed
	analytics.SetDB(dbPool)

	app.Get("/market", controllers.GetSummaryMarket)
	app.Get("/analytics/most-expensive-sale", controllers.GetMostExpensiveSale)
	app.Get("/analytics/average-delivery-time", controllers.GetAverageDeliveryTime)

	app.Post("/auth/register", controllers.Register)

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}
