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

	// Import a file
	file := []byte("YOUR_FILE_CONTENT")
	fileImportReq := model.FileImportRequest{
		File:     file,
		FileType: "YOUR_FILE_TYPE",
		DataType: 1,
		FormatID: 123,
		Encoding: "YOUR_FILE_ENCODING",
	}
	fileImportResp, err := handler.ImportFile(c.Token, fileImportReq)
	if err != nil {
		log.Println(err)
		return
	}
	if fileImportResp.Status != "OK" {
		log.Println(fileImportResp.ReturnValue)
		return
	}
}
