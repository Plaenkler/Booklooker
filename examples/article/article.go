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

	// Delete an article
	req := model.ArticleRequest{
		OrderNo: "YOUR_ORDER_NUMBER",
	}

	articleResp, err := handler.DeleteArticle(c.Token, req)
	if err != nil {
		log.Fatalf("failed to delete article: %v", err)
	}
	log.Println("Status:", articleResp.Status)
	log.Println("Return:", articleResp.ReturnValue)
}
