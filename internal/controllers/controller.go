package controllers

import (
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/analytics"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/models"
	"github.com/gofiber/fiber/v2"
)

// GetSummaryMarket godoc
// @Summary Get market summary
// @Description get market summary data
// @Tags market
// @Accept  json
// @Produce  json
// @Success 200 {string} string "GetSummaryMarket"
// @Failure 500 {string} string "Error"
// @Router /market [get]
func GetSummaryMarket(c *fiber.Ctx) error {
	marketSummary, err := analytics.GetSummaryMarket()
	if err != nil {
		return c.Status(500).SendString("Error")
	}
	return c.Status(200).JSON(marketSummary)
}

// GetMostExpensiveSale godoc
// @Summary Get most expensive sale
// @Description get most expensive sale data
// @Tags analytics
// @Accept  json
// @Produce  json
// @Success 200 {string} string "GetMostExpensiveSale"
// @Failure 500 {string} string "Error"
// @Router /analytics/most-expensive-sale [get]
func GetMostExpensiveSale(c *fiber.Ctx) error {
	return c.SendString("GetMostExpensiveSale")
}

// GetAverageDeliveryTime godoc
// @Summary Get average delivery time
// @Description get average delivery time data
// @Tags analytics
// @Accept  json
// @Produce  json
// @Success 200 {string} string "GetAverageDeliveryTime"
// @Failure 500 {string} string "Error"
// @Router /analytics/average-delivery-time [get]
func GetAverageDeliveryTime(c *fiber.Ctx) error {
	return c.SendString("GetAverageDeliveryTime")
}

// Register godoc
// @Summary Register
// @Description register a new user
// @Tags auth
// @Param Register body models.Register true "User registration details"
// @Accept  json
// @Produce  json
// @Success 201 {string} string "User added"
// @Failure 500 {string} string "Bad server"
// @Failure 400 {string} string "Bad request"
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var register models.Register
	if err := c.BodyParser(&register); err != nil {
		return c.Status(400).SendString("Bad request")
	}

	if err := analytics.AddUser(register); err != nil {
		return c.Status(500).SendString("Bad server")
	}

	return c.Status(201).SendString("User added")
}

// Login godoc
// @Summary Login
// @Description login a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param Login body models.Login true "User login details"
// @Success 200 {string} string "JWT"
// @Failure 500 {string} string "Bad server"
// @Failure 400 {string} string "Bad request"
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var login models.Login
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).SendString("Bad request")
	}

	if err := analytics.Login(login); err != nil {
		return c.Status(500).SendString("Bad server")
	}

	return c.SendString("JWT")
}
