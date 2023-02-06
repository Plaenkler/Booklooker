package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/plaenkler/booklooker/model"
)

func DeleteArticle(token model.Token, req model.ArticleRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	if req.OrderNo == "" {
		return nil, fmt.Errorf("orderNo is required")
	}
	url := model.BaseURL + model.ArticlePath + "?token=" + token.Value + "&orderNo=" + req.OrderNo
	httpReq, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var articleResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &articleResp)
	if err != nil {
		return nil, err
	}
	return &articleResp, nil
}
