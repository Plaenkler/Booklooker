package main

import (
	"encoding/json"
	"log"

	"github.com/plaenkler/booklooker/client"
	"github.com/plaenkler/booklooker/handler"
	"github.com/plaenkler/booklooker/model"
)

func main() {
	// Create a new client
	c := client.Client{APIKey: "YOUR_API_KEY"}
	err := c.Start()
	if err != nil {
		log.Fatalf("failed to start client: %v", err)
	}
	defer c.Stop()

	// Get orders for a specific date or time range
	req := model.OrderRequest{
		OrderID:  "123",        // Can only contain numbers
		Date:     "2023-02-01", // Will override DateFrom and DateTo
		DateFrom: "2023-02-01",
		DateTo:   "2023-02-02",
	}

	orderResp, err := handler.GetOrder(c.Token, req)
	if err != nil {
		log.Printf("error getting order: %v", err)
		return
	}
	if orderResp.Status != "OK" {
		log.Fatalf("error getting orders: %v", orderResp.ReturnValue)
	}
	var orders []model.Order
	err = json.Unmarshal([]byte(orderResp.ReturnValue), &orders)
	if err != nil {
		log.Println("error unmarshalling orders:", err)
	}
	for _, order := range orders {
		log.Println(order)
	}
}
