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
