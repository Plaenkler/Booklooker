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

func PutOrderMessage(token model.Token, req model.OrderMessageRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	params := url.Values{}
	if req.OrderID == "" {
		return nil, fmt.Errorf("orderId is not valid")
	}
	if req.MessageType == "" {
		return nil, fmt.Errorf("messageType is not set")
	}
	if req.AdditionalText != "" {
		params.Set("additionalText", req.AdditionalText)
	}
	params.Set("orderId", req.OrderID)
	params.Set("messageType", string(req.MessageType))
	url := model.BaseURL + model.OrderMessagePath + "?token=" + token.Value + "&" + params.Encode()
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
	var orderMessageResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderMessageResp)
	if err != nil {
		return nil, err
	}
	return &orderMessageResp, nil
}
