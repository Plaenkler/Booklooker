package main

import (
	"fmt"

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
		fmt.Println(err)
		return
	}
	if authResp.Status != "success" {
		fmt.Println(authResp.Message)
		return
	}
	token := authResp.Token

	//
	req := models.ArticleStatusRequest{OrderNo: "order_no_value"}

	resp, err := handler.GetArticleStatus(token, req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("response:", resp)
}
