package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func CancelOrder(token model.Token, req model.OrderCancelRequest) (*model.GlobalResponse, error) {
	url := model.BaseURL + model.OrderCancelPath + "?token=" + token.Value
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
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
