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

	// Cancel an order item
	req := &models.OrderItemCancelRequest{
		OrderItemId: "order_item_id",
		MediaType:   1,
	}
	orderItemCancelResp, err := handler.PutOrderItemCancel(token, req)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	if orderItemCancelResp.Status != "success" {
		log.Println(orderItemCancelResp.Message)
		return
	}
}
