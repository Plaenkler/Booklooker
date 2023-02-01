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

	// Cancel an order
	req := model.OrderCancelRequest{
		OrderID: "123",
	}

	orderCancelResp, err := handler.CancelOrder(c.Token, req)
	if err != nil {
		log.Println("Error cancelling order:", err)
		return
	}

	log.Println("Order cancel response:", orderCancelResp)
}
