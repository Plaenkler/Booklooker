package main

import (
	"log"

	"github.com/plaenkler/booklooker/api/handler"
	"github.com/plaenkler/booklooker/api/models"
)

func main() {
	req := models.AuthenticateRequest{APIKey: "YOUR_API_KEY"}
	authResp, err := handler.Authenticate(req)
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}
	log.Println("Authentication response:", authResp)
}
