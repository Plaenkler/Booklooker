package main

import (
	"log"
	"os"

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
	file, err := os.Open("export.txt")
	if err != nil {
		log.Println(err)
		return
	}
	fileImportReq := model.FileImportRequest{
		File:      file,        // txt or zip file
		FileType:  "article",   // or "pic" for picture(s)
		DataType:  1,           // 0: Add, change, delete, 1: Replace, 2: Delete
		MediaType: model.Books, // Possible values: 0: Books, 1: Movies, 2: Music, 3: Audio books, 4: Games
		// FormatID:  "BOOKLOOKER_FORMAT_ID", // Possible values not documented (expect n/a)
		// Encoding:  "YOUR_FILE_ENCODING",   // Default ISO8859-1 / Latin1 (n/a), IBMPC/CR (CP437), macintosh (Mac OS Roman), UTF-8
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
