package models

type Sale struct {
	Id              int     `json:"id"`
	Customer_email  string  `json:"customer_email"`
	Delivery_status string  `json:"delivery_status"`
	Total           float64 `json:"total"`
}

type MarketSummary struct {
	MarketName          string             `json:"market_name"`
	SalesMetric         map[string]float64 `json:"sales_metric"`
	MostExpensiveSale   Sale               `json:"most_expensive_sale"`
	AverageDeliveryTime float64            `json:"average_delivery_time"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username" validate:"required" example:"johndoe"`
	Email    string `json:"email" validate:"required,email" example:"example@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}
