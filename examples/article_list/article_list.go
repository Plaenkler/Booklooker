package main

import (
	"fmt"
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
		fmt.Println(err)
		return
	}
	if authResp.Status != "success" {
		fmt.Println(authResp.Message)
		return
	}
	token := authResp.Token

	//
	req := models.ArticleListRequest{
		Field:     "title",
		ShowPrice: true,
		ShowStock: false,
	}

	resp, err := handler.GetArticleList(token, req)
	if err != nil {
		log.Fatalf("failed to get article list: %v", err)
	}
	fmt.Println("Article List:", resp.ArticleList)
}
