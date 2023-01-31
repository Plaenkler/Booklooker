package main

import (
	"log"

	"github.com/plaenkler/booklooker/handler"
	"github.com/plaenkler/booklooker/model"
)

func main() {
	// If you don't want to use the client, you can authenticate directly
	req := model.AuthenticateRequest{APIKey: "YOUR_API_KEY"}
	authResp, err := handler.Authenticate(req)
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}
	log.Println("Status:", authResp.Status)
	log.Println("Return:", authResp.ReturnValue)
}
