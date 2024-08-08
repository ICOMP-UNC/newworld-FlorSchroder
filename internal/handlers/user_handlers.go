package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/models"
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/services"
	"github.com/gofiber/fiber/v2"
)

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

	if err := services.AddUser(register); err != nil {
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

	token, err := services.Login(login)
	if err != nil {
		return c.Status(500).SendString("Bad server")
	}

	return c.SendString(token)
}

// GetOffers godoc
// @Summary Get offers
// @Description get all offers
// @Tags auth
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Success 200 {string} string "JWT"
// @Failure 500 {string} string "Bad server"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/offers [get]
func GetOffers(c *fiber.Ctx) error {
	resp, err := http.Get("http://localhost:8080/supplies")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting supplies",
		})
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error reading response body",
		})
	}

	// Parse JSON
	var offers map[string]models.OfferWithID
	if err := json.Unmarshal(body, &offers); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error unmarshaling JSON",
		})
	}

	// Add offers to database
	for _, offer := range offers {
		if err := services.AddOffer(offer); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error adding offer to database",
			})
		}
	}

	jwtToken := c.Get("Authorization")
	//println("GetOffers " + jwtToken)
	if jwtToken == "" {
		return c.Status(401).JSON(models.Message{Status: "Unauthorized"})
	}

	//println("GetOffers " + jwtToken)

	// Get offers from database
	offers2, err := services.GetOffers(jwtToken)
	if err != nil {
		return c.Status(500).JSON(models.Message{Status: "Bad server"})
	}

	return c.Status(200).JSON(offers2)
}

// Checkout godoc
// @Summary Checkout
// @Description checkout
// @Tags auth
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Param Order body models.Order true "Order details"
// @Success 200 {object} models.Message "Order placed"
// @Failure 500 {string} string "Bad server"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/checkout [post]
func Checkout(c *fiber.Ctx) error {
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(401).JSON(models.Message{Status: "Unauthorized"})
	}

	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(models.Message{Status: "Bad request"})
	}

	message, err := services.Checkout(order, jwtToken)
	if err != nil {
		return c.Status(500).JSON(models.Message{Status: "Bad server"})
	}

	return c.Status(200).JSON(message)
}

// GetDashboard godoc
// @Summary Get dashboard
// @Description get dashboard data
// @Tags admin
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT"
// @Success 200 {string} string "Dashboard data"
// @Failure 500 {string} string "Bad server"
// @Failure 401 {string} string "Unauthorized"
// @Router /admin/dashboard [get]
func GetDashboard(c *fiber.Ctx) error {
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(401).SendString("Unauthorized")
	}

	dashboard, err := services.GetDashboard(jwtToken)
	if err != nil {
		return c.Status(500).SendString("Bad server")
	}
	return c.Status(200).JSON(dashboard)
}

// getOrderStatus godoc
// @Summary Get order status
// @Description get order status
// @Tags auth
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param Authorization header string true "JWT"
// @Success 200 {string} string "Order status"
// @Failure 500 {string} string "Bad server"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/orders/{id} [get]
func GetOrderStatus(c *fiber.Ctx) error {
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(401).SendString("Unauthorized")
	}

	orderID := c.Params("id")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return c.Status(400).SendString("Invalid order ID")
	}

	status, err := services.GetOrderStatus(id, jwtToken)
	if err != nil {
		return c.Status(500).SendString("Bad server")
	}
	return c.Status(200).SendString(status)
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description update order status
// @Tags auth
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param Authorization header string true "JWT"
// @Param status body string true "Order status"  example:"preparing/processing/shipped/delivered"
// @Success 200 {string} string "Order status updated"
// @Failure 500 {string} string "Bad server"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/order/{id} [patch]
func UpdateOrderStatus(c *fiber.Ctx) error {
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(401).SendString("Unauthorized")
	}

	orderID := c.Params("id")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return c.Status(400).SendString("Invalid order ID")
	}

	var status string
	if err := c.BodyParser(&status); err != nil {
		return c.Status(400).SendString("Bad request")
	}

	err = services.UpdateOrderStatus(id, status, jwtToken)
	if err != nil {
		return c.Status(500).SendString("Bad server")
	}

	return c.Status(200).SendString("Order status updated")
}
