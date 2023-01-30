package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func DeleteArticle(token string, req models.ArticleRequest) (*models.GlobalResponse, error) {
	url := baseURL + models.ArticlePath + "?token=" + token + "&orderNo=" + req.OrderNo
	httpReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
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
	var articleResp models.GlobalResponse
	err = json.Unmarshal(jsonResp, &articleResp)
	if err != nil {
		return nil, err
	}
	return &articleResp, nil
}
