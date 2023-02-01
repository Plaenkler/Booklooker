package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/plaenkler/booklooker/model"
)

func GetOrder(token model.Token, req model.OrderRequest) (*model.OrderResponse, error) {
	params := url.Values{}
	if reflect.ValueOf(req.OrderID).IsZero() {
		return nil, fmt.Errorf("orderId is not set")
	}
	if !reflect.ValueOf(req.Date).IsZero() {
		params.Set("date", req.Date)
	}
	if !params.Has("date") {
		if reflect.ValueOf(req.DateFrom).IsZero() || reflect.ValueOf(req.DateTo).IsZero() {
			return nil, fmt.Errorf("time range not set")
		}
		params.Set("dateFrom", req.DateFrom)
		params.Set("dateTo", req.DateTo)
	}
	url := model.BaseURL + model.OrderPath + "?token=" + token.Value + "&" + params.Encode()
	resp, err := http.Get(url)
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
