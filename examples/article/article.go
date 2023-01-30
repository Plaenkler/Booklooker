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
	req := models.ArticleRequest{
		OrderNo: "123",
	}

	resp, err := handler.DeleteArticle(token, req)
	if err != nil {
		log.Fatalf("failed to delete article: %v", err)
	}
	fmt.Println("Article Deletion Status:", resp.Status)
}
