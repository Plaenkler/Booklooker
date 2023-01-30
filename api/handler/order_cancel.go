package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func CancelOrder(token string, req models.OrderCancelRequest) (*models.OrderCancelResponse, error) {
	url := baseURL + models.OrderCancelPath + "?token=" + token
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
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
	var orderCancelResp models.OrderCancelResponse
	err = json.Unmarshal(jsonResp, &orderCancelResp)
	if err != nil {
		return nil, err
	}
	return &orderCancelResp, nil
}
