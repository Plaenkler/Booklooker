package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func GetArticleStatus(token string, req models.ArticleStatusRequest) (*models.GlobalResponse, error) {
	url := baseURL + models.ArticleStatusPath + "?token=" + token + "&orderNo=" + req.OrderNo
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var articleStatusResp models.GlobalResponse
	err = json.Unmarshal(jsonResp, &articleStatusResp)
	if err != nil {
		return nil, err
	}
	return &articleStatusResp, nil
}
