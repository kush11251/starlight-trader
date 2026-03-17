package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"starlight-trader/src/models"
	"starlight-trader/src/controllers"
)

func main() {
	router := mux.NewRouter()
	controllers.RegisterRoutes(router)
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(:8000, router)
}
