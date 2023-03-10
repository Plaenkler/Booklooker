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

	// Set the order status
	req := model.OrderStatusRequest{
		OrderID: "123",
		Status:  "CANCELED", // See possible return values in the documentation
	}
	orderStatusResp, err := handler.SetOrderStatus(c.Token, req)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Status:", orderStatusResp.Status)
	log.Println("Return:", orderStatusResp.ReturnValue)
}
