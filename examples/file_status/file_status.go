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

	// Get file status
	fileStatusResp, err := handler.GetFileStatus(token, models.FileStatusRequest{Filename: "your_filename"})
	if err != nil {
		log.Fatalf("error getting file status: %v", err)
	}
	log.Println("file status:", fileStatusResp.Status)
}
