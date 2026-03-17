package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"starlight-trader/src/models"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc(/orders, getOrders).Methods(http.MethodGet)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	orders := []models.Order{{ID: "1", Symbol: "AAPL", Quantity: 10, Price: 150.0}}
	json.NewEncoder(w).Encode(orders)
}
