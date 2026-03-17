package utils

import (
	"math"
)

func CalculateRisk(order models.Order) float64 {
	// Simplified risk calculation example
	return order.Price * float64(order.Quantity) * 0.1
}
