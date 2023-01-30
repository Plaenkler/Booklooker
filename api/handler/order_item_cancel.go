package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func PutOrderItemCancel(token string, req *models.OrderItemCancelRequest) (*models.GlobalResponse, error) {
	url := baseURL + models.OrderItemCancelPath + "?token=" + token
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	httpReq, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var orderItemCancelResp models.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderItemCancelResp)
	if err != nil {
		return nil, err
	}
	return &orderItemCancelResp, nil
}
