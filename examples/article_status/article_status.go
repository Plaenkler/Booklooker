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

	// Query the status of an item
	req := model.ArticleStatusRequest{OrderNo: "YOUR_ORDER_NUMBER"}
	articleStatusResp, err := handler.GetArticleStatus(c.Token, req)
	if err != nil {
		log.Println("error:", err)
		return
	}
	log.Println("Status:", articleStatusResp.Status)
	log.Println("Return:", articleStatusResp.ReturnValue)
}
