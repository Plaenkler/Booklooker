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
		log.Println("Status:", authResp.Status)
		log.Println("Return:", authResp.ReturnValue)
		return
	}
	token := authResp.ReturnValue
	log.Println("Token:", token)

	// Delete an article
	req := models.ArticleRequest{
		OrderNo: "123",
	}

	articleResp, err := handler.DeleteArticle(token, req)
	if err != nil {
		log.Fatalf("failed to delete article: %v", err)
	}
	log.Println("Status:", articleResp.Status)
	log.Println("Return:", articleResp.ReturnValue)
}
