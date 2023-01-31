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

	// Import a file
	file := []byte("your file content")
	fileImportReq := model.FileImportRequest{
		File:     file,
		FileType: "your file type",
		DataType: 1,
		FormatID: 123,
		Encoding: "your encoding",
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
