package main

import (
	"log"

	"github.com/plaenkler/booklooker/client"
	"github.com/plaenkler/booklooker/handler"
	"github.com/plaenkler/booklooker/model"
)

func main() {
	// Create a new client
	c := client.Client{APIKey: "YOUR_API_KEY"}
	err := c.Start()
	if err != nil {
		log.Fatalf("failed to start client: %v", err)
	}
	defer c.Stop()

	// Cancel an order item
	req := &model.OrderItemCancelRequest{
		OrderItemID: "order_item_id",
		MediaType:   1,
	}
	orderItemCancelResp, err := handler.PutOrderItemCancel(c.Token, req)
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
