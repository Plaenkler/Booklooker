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
		log.Println("Status:", authResp.Status)
		log.Println("Return:", authResp.ReturnValue)
		return
	}
	token := authResp.ReturnValue
	log.Println("Token:", token)

	// Get orders for a specific date or time range
	req := models.OrderRequest{
		OrderID:  "123",
		Date:     "2022-12-31",
		DateFrom: "",
		DateTo:   "",
	}

	orderResp, err := handler.GetOrder(token, req)
	if err != nil {
		log.Println("Error getting order:", err)
		return
	}

	log.Println("Order response:", orderResp)
}
