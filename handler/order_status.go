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

func SetOrderStatus(token model.Token, req model.OrderStatusRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	params := url.Values{}
	if req.OrderID == "" {
		return nil, fmt.Errorf("orderId is not set")
	}
	if req.Status == "" {
		return nil, fmt.Errorf("status is not set")
	}
	params.Set("orderId", req.OrderID)
	params.Set("status", req.Status)
	url := model.BaseURL + model.OrderStatusPath + "?token=" + token.Value + "&" + params.Encode()
	httpReq, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}
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
	var orderStatusResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderStatusResp)
	if err != nil {
		return nil, err
	}
	return &orderStatusResp, nil
}
