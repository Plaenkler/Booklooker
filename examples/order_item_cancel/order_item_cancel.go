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

	// Cancel an order item
	req := &models.OrderItemCancelRequest{
		OrderItemId: "order_item_id",
		MediaType:   1,
	}
	orderItemCancelResp, err := handler.PutOrderItemCancel(token, req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if orderItemCancelResp.Status != "success" {
		fmt.Println(orderItemCancelResp.Message)
		return
	}
}
