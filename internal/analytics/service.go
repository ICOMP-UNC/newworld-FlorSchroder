package analytics

import "github.com/ICOMP-UNC/newworld-FlorSchroder/internal/models"

func GetSummaryMarket() (models.MarketSummary, error) {
	return models.MarketSummary{
		MarketName: "Market",
		SalesMetric: map[string]float64{
			"total_sales": 1000,
		},
		MostExpensiveSale: models.Sale{
			Id:              1,
			Customer_email:  "example@example.com",
			Delivery_status: "delivered",
			Total:           1000,
		},
		AverageDeliveryTime: 10,
	}, nil
}
