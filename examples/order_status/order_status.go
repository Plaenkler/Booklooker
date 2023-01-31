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
	c.Start()
	defer c.Stop()

	// Set the order status
	req := model.OrderStatusRequest{
		OrderID: "123",
		Status:  "complete",
	}
	orderStatusResp, err := handler.SetOrderStatus(c.Token, req)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Response:", orderStatusResp)
}
