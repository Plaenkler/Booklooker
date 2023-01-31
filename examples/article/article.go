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
	c.Start()
	defer c.Stop()

	// Delete an article
	req := model.ArticleRequest{
		OrderNo: "Example123",
	}

	articleResp, err := handler.DeleteArticle(c.Token, req)
	if err != nil {
		log.Fatalf("failed to delete article: %v", err)
	}
	log.Println("Status:", articleResp.Status)
	log.Println("Return:", articleResp.ReturnValue)
}
