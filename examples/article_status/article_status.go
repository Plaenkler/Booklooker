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

	// Query the status of an item
	req := models.ArticleStatusRequest{OrderNo: "Order123"}
	articleStatusResp, err := handler.GetArticleStatus(token, req)
	if err != nil {
		log.Println("error:", err)
		return
	}
	log.Println("Status:", articleStatusResp.Status)
	log.Println("Return:", articleStatusResp.ReturnValue)
}
