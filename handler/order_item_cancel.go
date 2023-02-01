package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/plaenkler/booklooker/model"
)

func PutOrderItemCancel(token model.Token, req model.OrderItemCancelRequest) (*model.GlobalResponse, error) {
	params := url.Values{}
	if reflect.ValueOf(req.MediaType).IsZero() {
		return nil, fmt.Errorf("mediaType is not valid")
	}
	if reflect.ValueOf(req.OrderItemID).IsZero() {
		return nil, fmt.Errorf("orderItemId is not set")
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
