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

	// Send a message to the customer
	req := model.OrderMessageRequest{
		OrderID:        "123", // Can only contain numbers
		MessageType:    model.ShippingNotice,
		AdditionalText: "YOUR_ADDITIONAL_TEXT",
	}
	orderMessageResp, err := handler.PutOrderMessage(c.Token, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Status:", orderMessageResp.Status)
	log.Println("Return:", orderMessageResp.ReturnValue)
}
