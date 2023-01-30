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

	// Set the order status
	req := models.OrderStatusRequest{
		OrderID: "123",
		Status:  "complete",
	}
	resp, err := handler.SetOrderStatus(token, req)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Response:", resp)
}
