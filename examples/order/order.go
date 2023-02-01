package main

import (
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
		OrderID:  "123",
		Date:     "2022-12-31",
		DateFrom: "",
		DateTo:   "",
	}

	orderResp, err := handler.GetOrder(c.Token, req)
	if err != nil {
		log.Println("Error getting order:", err)
		return
	}

	log.Println("Order response:", orderResp)
}
