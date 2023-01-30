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

	// Download all active article numbers
	req := models.ArticleListRequest{
		ReturnType: "isbn", // Possible values: orderNo, ISBN or EAN
		ShowPrice:  true,
		ShowStock:  true,
		MediaType:  0, // Possible values: 0: Books, 1: Movies, 2: Music, 3: Audio books, 4: Games or n/a
	}

	articleListResp, err := handler.GetArticleList(token, req)
	if err != nil {
		log.Fatalf("failed to get article list: %v", err)
	}
	log.Println("Status:", articleListResp.Status)
	log.Println("Return:", articleListResp.ReturnValue)
}
