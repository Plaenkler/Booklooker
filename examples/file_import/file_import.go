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

	// Import a file
	file := []byte("your file content")
	fileImportReq := models.FileImportRequest{
		File:     file,
		FileType: "your file type",
		DataType: 1,
		FormatID: 123,
		Encoding: "your encoding",
	}
	fileImportResp, err := handler.ImportFile(token, fileImportReq)
	if err != nil {
		log.Println(err)
		return
	}
	if fileImportResp.Status != "success" {
		log.Println(fileImportResp.Message)
		return
	}
}