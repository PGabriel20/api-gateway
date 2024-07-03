package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
    ID     int    `json:"id"`
    Item   string `json:"item"`
    Price int    `json:"price"`
}

func getOrders(w http.ResponseWriter, r *http.Request) {
    orders := []Order{
        {ID: 1, Item: "Macbook", Price: 1000},
        {ID: 2, Item: "Phone", Price: 500},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}

func main() {
    http.HandleFunc("/orders", getOrders)
    log.Println("Order Service running on port 8082")
    log.Fatal(http.ListenAndServe(":8082", nil))
}
