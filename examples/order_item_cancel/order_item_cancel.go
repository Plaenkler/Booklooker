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

	// Cancel an order item
	req := &models.OrderItemCancelRequest{
		OrderItemID: "order_item_id",
		MediaType:   1,
	}
	orderItemCancelResp, err := handler.PutOrderItemCancel(token, req)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	if orderItemCancelResp.Status != "OK" {
		log.Println(orderItemCancelResp.ReturnValue)
		return
	}
	log.Println("Status:", orderItemCancelResp.Status)
	log.Println("Return:", orderItemCancelResp.ReturnValue)
}
