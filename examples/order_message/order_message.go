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

	// Send a message to the customer
	req := model.OrderMessageRequest{
		OrderID:        "ORDER_ID",
		MessageType:    "MESSAGE_TYPE",
		AdditionalText: "ADDITIONAL_TEXT",
	}
	orderMessageResp, err := handler.PutOrderMessage(c.Token, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Status:", orderMessageResp.Status)
	log.Println("Return:", orderMessageResp.ReturnValue)
}
