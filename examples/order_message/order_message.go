package main

import (
	"fmt"
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
		fmt.Println(err)
		return
	}
	if authResp.Status != "success" {
		fmt.Println(authResp.Message)
		return
	}
	token := authResp.Token

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
	fmt.Println("Status:", resp.Status)
	fmt.Println("Message:", resp.Message)
}
