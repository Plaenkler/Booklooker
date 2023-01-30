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

	//
	req := models.ArticleRequest{
		OrderNo: "123",
	}

	resp, err := handler.DeleteArticle(token, req)
	if err != nil {
		log.Fatalf("failed to delete article: %v", err)
	}
	log.Println("Article Deletion Status:", resp.Status)
}
