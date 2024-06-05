package main

import (
	controller "github.com/ICOMP-UNC/newworld-FlorSchroder/internal/controllers"

	_ "github.com/ICOMP-UNC/newworld-FlorSchroder/docs" // Ensure this path is correct

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
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

	app.Get("/market", controller.GetSummaryMarket)
	app.Get("/analytics/most-expensive-sale", controller.GetMostExpensiveSale)
	app.Get("/analytics/average-delivery-time", controller.GetAverageDeliveryTime)

	app.Get("/swagger/*", swagger.HandlerDefault) // Serve the Swagger documentation

	app.Listen(":3000")
}

// package main

// import (
// 	controller "goApp/cmd/controllers"

// 	"github.com/gofiber/fiber/v2"

// 	"github.com/gofiber/swagger"

// 	_ "goApp/cmd/server/docs"
// )

// // @title Fiber API Example
// // @version 1.0
// // @description This is a sample server.
// // @host localhost:3000
// // @BasePath /
// // @termsOfService http://swagger.io/terms/
// // @contact.name API Support
// // @contact.email you@example.com
// // @license.name Apache 2.0
// // @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// func main() {
// 	app := fiber.New()

// 	app.Get("/market", controller.GetSummaryMarket)

// 	app.Get("/swagger/*", swagger.HandlerDefault)

// 	// app.Get("/swagger/*", swagger.New(swagger.Config{
// 	// 	URL:         "/swagger/doc.json",
// 	// 	DeepLinking: false,
// 	// }))

// 	app.Listen(":3000")
// }
