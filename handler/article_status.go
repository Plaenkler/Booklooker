package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func GetArticleStatus(token model.Token, req model.ArticleStatusRequest) (*model.GlobalResponse, error) {
	if req.OrderNo == "" {
		return nil, fmt.Errorf("orderNo is required")
	}
	url := model.BaseURL + model.ArticleStatusPath + "?token=" + token.Value + "&orderNo=" + req.OrderNo
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var articleStatusResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &articleStatusResp)
	if err != nil {
		return nil, err
	}
	return &articleStatusResp, nil
}
