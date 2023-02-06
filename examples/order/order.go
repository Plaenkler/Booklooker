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
		Date: "2023-02-05", // Will override DateFrom and DateTo
		// DateFrom: "2021-01-01",
		// DateTo:   "2021-01-31",
		// OrderID:  "123456789",
	}

	orderResp, err := handler.GetOrder(c.Token, req)
	if err != nil {
		log.Printf("error getting order: %v", err)
		return
	}
	log.Printf("Status: %v", orderResp.Status)
	log.Printf("Return: %+v", orderResp.ReturnValue)
}
