package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func Authenticate(req models.AuthenticateRequest) (*models.GeneralResponse, error) {
	url := baseURL + models.AuthenticatePath + "?apiKey=" + req.APIKey
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var authResp models.GeneralResponse
	err = json.Unmarshal(jsonResp, &authResp)
	if err != nil {
		return nil, err
	}
	return &authResp, nil
}
