package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/plaenkler/booklooker/model"
)

func GetArticleList(token model.Token, req model.ArticleListRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	params := url.Values{}
	if req.MediaType != "" {
		params.Set("mediaType", string(req.MediaType))

		// Can only be used if mediaType is present
		if req.ReturnType != "" {
			params.Set("field", req.ReturnType)
		}
	}
	if req.ShowPrice {
		params.Set("showPrice", "1")
	}
	if req.ShowStock {
		params.Set("showStock", "1")
	}
	url := model.BaseURL + model.ArticleListPath + "?token=" + token.Value + "&" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var articleListResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &articleListResp)
	if err != nil {
		return nil, err
	}
	return &articleListResp, nil
}
