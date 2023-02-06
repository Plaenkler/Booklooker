package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/plaenkler/booklooker/model"
)

func CancelOrder(token model.Token, req model.OrderCancelRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	if req.OrderID == "" {
		return nil, fmt.Errorf("orderID is required")
	}
	url := model.BaseURL + model.OrderCancelPath + "?token=" + token.Value + "&orderId=" + req.OrderID
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
	var orderCancelResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderCancelResp)
	if err != nil {
		return nil, err
	}
	return &orderCancelResp, nil
}
