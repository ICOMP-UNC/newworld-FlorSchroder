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

type Login struct {
	Username string `json:"username" validate:"required" example:"johndoe"`
	Email    string `json:"email" validate:"required,email" example:"example@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type Offer struct {
	Name     string  `json:"name" validate:"required" example:"meat"`     // Name of the offer
	Quantity int     `json:"quantity" validate:"required" example:"10"`   // Quantity of the offer
	Category string  `json:"category" validate:"required" example:"food"` // Category of the offer
	Price    float64 `json:"price" validate:"required" example:"10.5"`    // Price of the offer
}

type OfferWithID struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type OrderItem struct {
	Quantity  int `json:"quantity" example:"10" validate:"required"`
	ProductID int `json:"product_id" example:"1" validate:"required"`
}

type Order struct {
	Items []OrderItem `json:"items" example:"[{\"quantity\": 10, \"product_id\": 1}, {\"quantity\": 5, \"product_id\": 2}]" validate:"required"`
}

// type OrderItem struct {
// 	Quantity  int `json:"quantity" validate:"required" example:"10"`
// 	ProductID int `json:"product_id" validate:"required" example:"1"`
// }

// type Order struct {
// 	// ID     int         `json:"id" `
// 	Items []OrderItem `json:"items" validate:"required" example:"[{\"quantity\": 10, \"product_id\": 1} ,{\"quantity\": 5, \"product_id\": 2}]"`
// 	// Status string      `json:"status" validate:"required" example:"preparing"` //"preparing/processing/shipped/delivered"
// }

type Message struct {
	Total  float64 `json:"total"`
	Status string  `json:"status"`
}

type Dashboard struct {
	Orders []Order
}
