package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func DeleteArticle(token string, req models.ArticleRequest) (*models.ArticleResponse, error) {
	url := baseURL + models.ArticlePath + "?token=" + token
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonReq))
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
	var articleResp models.ArticleResponse
	err = json.Unmarshal(jsonResp, &articleResp)
	if err != nil {
		return nil, err
	}
	return &articleResp, nil
}
