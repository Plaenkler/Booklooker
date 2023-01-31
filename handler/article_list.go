package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/plaenkler/booklooker/model"
)

func GetArticleList(token model.Token, req model.ArticleListRequest) (*model.GlobalResponse, error) {
	params := url.Values{}
	if !reflect.ValueOf(req.MediaType).IsZero() {
		params.Set("mediaType", string(req.MediaType))

		// Can only be used if mediaType is present
		if !reflect.ValueOf(req.ReturnType).IsZero() {
			params.Set("field", req.ReturnType)
		}
	}
	if !reflect.ValueOf(req.ShowPrice).IsZero() {
		params.Set("showPrice", "1")
	}
	if !reflect.ValueOf(req.ShowStock).IsZero() {
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