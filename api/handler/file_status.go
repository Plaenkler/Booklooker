package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func GetFileStatus(token string, req models.FileStatusRequest) (*models.FileStatusResponse, error) {
	url := baseURL + models.FileStatusPath + "?token=" + token + "&filename=" + req.Filename
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var fileStatusResp models.FileStatusResponse
	err = json.Unmarshal(jsonResp, &fileStatusResp)
	if err != nil {
		return nil, err
	}
	return &fileStatusResp, nil
}
