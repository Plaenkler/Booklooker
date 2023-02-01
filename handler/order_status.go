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

func SetOrderStatus(token model.Token, req model.OrderStatusRequest) (*model.GlobalResponse, error) {
	params := url.Values{}
	if reflect.ValueOf(req.OrderID).IsZero() {
		return nil, fmt.Errorf("orderId is not set")
	}
	if reflect.ValueOf(req.Status).IsZero() {
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
