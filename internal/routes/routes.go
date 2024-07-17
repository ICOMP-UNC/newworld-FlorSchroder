package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/handlers"
)

func InitRoutes(app *fiber.App, dbPool *pgxpool.Pool) {
	app.Get("/auth/offers", handlers.GetOffers)
	app.Get("/admin/dashboard", handlers.GetDashboard)
	app.Get("/auth/orders/:id", handlers.GetOrderStatus)

	app.Post("/auth/register", handlers.Register)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/offer", handlers.AddOffer)
	app.Post("/auth/checkout", handlers.Checkout)

	app.Patch("/auth/order/:id", handlers.UpdateOrderStatus)
}
