package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func PutOrderMessage(token model.Token, req model.OrderMessageRequest) (*model.GlobalResponse, error) {
	url := model.BaseURL + model.OrderMessagePath + "?token=" + token.Value
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
	var orderMessageResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderMessageResp)
	if err != nil {
		return nil, err
	}
	return &orderMessageResp, nil
}
