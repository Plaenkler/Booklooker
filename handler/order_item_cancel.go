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

func PutOrderItemCancel(token model.Token, req model.OrderItemCancelRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	params := url.Values{}
	if req.MediaType == "" {
		return nil, fmt.Errorf("mediaType is required")
	}
	if req.OrderItemID == "" {
		return nil, fmt.Errorf("orderItemId is required")
	}
	params.Set("mediaType", string(req.MediaType))
	params.Set("orderItemId", req.OrderItemID)
	url := model.BaseURL + model.OrderItemCancelPath + "?token=" + token.Value + "&" + params.Encode()
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
	var orderItemCancelResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &orderItemCancelResp)
	if err != nil {
		return nil, err
	}
	return &orderItemCancelResp, nil
}
