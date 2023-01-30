package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func SetOrderStatus(token string, req models.OrderStatusRequest) (*models.GlobalResponse, error) {
	url := baseURL + models.OrderStatusPath + "?token=" + token
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
	var orderStatusResp models.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderStatusResp)
	if err != nil {
		return nil, err
	}
	return &orderStatusResp, nil
}
