package controllers

import (
	"github.com/ICOMP-UNC/newworld-FlorSchroder/internal/analytics"

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

// package controllers

// import (
// 	"goApp/internal/analytics"

// 	"github.com/gofiber/fiber/v2"
// )

// // @Summary Get summary market
// // @Description Get summary market
// // @Tags Market
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} MarketSummary
// // @Router /market/summary [get]
// func GetSummaryMarket(c *fiber.Ctx) error {
// 	marketSummary, err := analytics.GetSummaryMarket()

// 	if err != nil {
// 		return c.Status(500).SendString("Error")
// 	}
// 	return c.Status(200).JSON(marketSummary)
// }

// // @Summary Get most expensive sale
// // @Description Get most expensive sale
// // @Tags Market
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} MostExpensiveSale
// // @Router /market/expensive [get]
// func GetMostExpensiveSale(c *fiber.Ctx) error {
// 	return c.SendString("GetMostExpensiveSale")
// }

// // @Summary Get average delivery time
// // @Description Get average delivery time
// // @Tags Market
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} AverageDeliveryTime
// // @Router /market/average [get]
// func GetAverageDeliveryTime(c *fiber.Ctx) error {
// 	return c.SendString("GetAverageDeliveryTime")
// }
