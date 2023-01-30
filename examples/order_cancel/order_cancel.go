package main

import (
	"fmt"

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

	// Cancel an order
	req := models.OrderCancelRequest{
		OrderID: "123",
	}

	orderCancelResp, err := handler.CancelOrder(token, req)
	if err != nil {
		fmt.Println("Error cancelling order:", err)
		return
	}

	fmt.Println("Order cancel response:", orderCancelResp)
}
