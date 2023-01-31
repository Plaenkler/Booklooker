package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func Authenticate(req model.AuthenticateRequest) (*model.GlobalResponse, error) {
	url := model.BaseURL + model.AuthenticatePath + "?apiKey=" + req.APIKey
	resp, err := http.Post(url, "application/json", nil)
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
