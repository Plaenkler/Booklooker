package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/plaenkler/booklooker/api/models"
)

func GetArticleList(token string, req models.ArticleListRequest) (*models.ArticleListResponse, error) {
	values := url.Values{}
	values.Set("field", req.Field)
	values.Set("showPrice", strconv.FormatBool(req.ShowPrice))
	values.Set("showStock", strconv.FormatBool(req.ShowStock))
	values.Set("mediaType", string(req.MediaType))

	url := baseURL + models.ArticleListPath + "?token=" + token + "&" + values.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var articleListResp models.ArticleListResponse
	err = json.Unmarshal(jsonResp, &articleListResp)
	if err != nil {
		return nil, err
	}
	return &articleListResp, nil
}
