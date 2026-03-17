package models

import (
	"encoding/json"
)

type Order struct {
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
	Quantity int    `json:"quantity"`
	Price    float64 `json:"price"`
}
