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

func GetOrder(token model.Token, req model.OrderRequest) (*model.OrderResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	params := url.Values{}
	if req.OrderID != "" {
		params.Set("orderID", req.OrderID)
	}
	if req.Date != "" {
		params.Set("date", req.Date)
	}
	if !params.Has("date") {
		if req.DateFrom == "" || req.DateTo == "" {
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
	var statusResp model.ErrorResponse
	err = json.Unmarshal(jsonResp, &statusResp)
	if err != nil {
		return nil, err
	}
	if statusResp.Status != "OK" {
		return nil, fmt.Errorf("NOK %v", statusResp.ReturnValue)
	}
	var orderResp model.OrderResponse
	err = json.Unmarshal(jsonResp, &orderResp)
	if err != nil {
		return nil, err
	}
	return &orderResp, nil
}
