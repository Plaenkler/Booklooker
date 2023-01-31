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
	c.Start()
	defer c.Stop()

	// Get file status
	fileStatusResp, err := handler.GetFileStatus(c.Token, model.FileStatusRequest{Filename: "your_filename"})
	if err != nil {
		log.Fatalf("error getting file status: %v", err)
	}
	log.Println("file status:", fileStatusResp.Status)
}
