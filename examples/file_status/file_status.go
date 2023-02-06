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

	// Get file status
	req := model.FileStatusRequest{Filename: "YOUR_FILE_NAME"}
	fileStatusResp, err := handler.GetFileStatus(c.Token, req)
	if err != nil {
		log.Fatalf("error getting file status: %v", err)
	}
	log.Println("Status:", fileStatusResp.Status)
	log.Println("Return:", fileStatusResp.ReturnValue)
}
