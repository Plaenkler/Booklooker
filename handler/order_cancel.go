package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func CancelOrder(token model.Token, req model.OrderCancelRequest) (*model.GlobalResponse, error) {
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
