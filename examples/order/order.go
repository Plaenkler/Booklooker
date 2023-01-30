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

	// Get orders for a specific date or time range
	req := models.OrderRequest{
		OrderID:  "123",
		Date:     "2022-12-31",
		DateFrom: "",
		DateTo:   "",
	}

	orderResp, err := handler.GetOrder(token, req)
	if err != nil {
		fmt.Println("Error getting order:", err)
		return
	}

	fmt.Println("Order response:", orderResp)
}
