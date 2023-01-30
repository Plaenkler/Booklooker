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

	// Cancel an order
	req := models.OrderCancelRequest{
		OrderID: "123",
	}

	orderCancelResp, err := handler.CancelOrder(token, req)
	if err != nil {
		log.Println("Error cancelling order:", err)
		return
	}

	log.Println("Order cancel response:", orderCancelResp)
}
