package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func Authenticate(req model.AuthenticateRequest) (*model.GlobalResponse, error) {
	if req.APIKey == "" {
		return nil, fmt.Errorf("apiKey is required")
	}
	url := model.BaseURL + model.AuthenticatePath + "?apiKey=" + req.APIKey
	resp, err := http.Post(url, "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var authResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &authResp)
	if err != nil {
		return nil, err
	}
	return &authResp, nil
}
