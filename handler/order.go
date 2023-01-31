package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func GetOrder(token model.Token, req model.OrderRequest) (*model.OrderResponse, error) {
	url := model.BaseURL + model.OrderPath + "?token=" + token.Value

	query := url + "&orderId=" + req.OrderID
	if req.Date != "" {
		query = query + "&date=" + req.Date
	}
	if req.DateFrom != "" {
		query = query + "&dateFrom=" + req.DateFrom
	}
	if req.DateTo != "" {
		query = query + "&dateTo=" + req.DateTo
	}

	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var orderResp model.OrderResponse
	err = json.Unmarshal(jsonResp, &orderResp)
	if err != nil {
		return nil, err
	}
	return &orderResp, nil
}
