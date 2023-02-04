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

	// Download all active article numbers
	req := model.ArticleListRequest{
		ReturnType: "isbn", // Possible values: orderNo, ISBN or EAN
		ShowPrice:  true,
		ShowStock:  true,
		MediaType:  model.Books, // Possible values: 0: Books, 1: Movies, 2: Music, 3: Audio books, 4: Games or n/a
	}

	articleListResp, err := handler.GetArticleList(c.Token, req)
	if err != nil {
		log.Fatalf("failed to get article list: %v", err)
	}
	log.Println("Status:", articleListResp.Status)
	log.Println("Return:", articleListResp.ReturnValue)
}
