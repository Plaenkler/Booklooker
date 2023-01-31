package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func SetOrderStatus(token model.Token, req model.OrderStatusRequest) (*model.GlobalResponse, error) {
	url := model.BaseURL + model.OrderStatusPath + "?token=" + token.Value
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
	var orderStatusResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderStatusResp)
	if err != nil {
		return nil, err
	}
	return &orderStatusResp, nil
}
