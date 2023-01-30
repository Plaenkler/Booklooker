package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func Authenticate(req models.AuthenticateRequest) (*models.AuthenticateResponse, error) {
	url := baseURL + models.AuthenticatePath
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var authResp models.AuthenticateResponse
	err = json.Unmarshal(jsonResp, &authResp)
	if err != nil {
		return nil, err
	}
	return &authResp, nil
}
