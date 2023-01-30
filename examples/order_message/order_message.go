package main

import (
	"log"

	"github.com/plaenkler/booklooker/api/handler"
	"github.com/plaenkler/booklooker/api/models"
)

func main() {
	// Authenticate to obtain a token
	authReq := models.AuthenticateRequest{
		APIKey: "your_api_key",
	}
	authResp, err := handler.Authenticate(authReq)
	if err != nil {
		log.Println(err)
		return
	}
	if authResp.Status != "OK" {
		log.Println(authResp.ReturnValue)
		return
	}
	token := authResp.ReturnValue

	// Send a message to the customer
	req := models.OrderMessageRequest{
		OrderID:        "ORDER_ID",
		MessageType:    "MESSAGE_TYPE",
		AdditionalText: "ADDITIONAL_TEXT",
	}
	resp, err := handler.PutOrderMessage(token, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Message:", resp.Message)
}
