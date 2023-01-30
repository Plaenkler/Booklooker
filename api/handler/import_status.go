package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func GetImportStatus(token string) (*models.GlobalResponse, error) {
	url := baseURL + models.ImportStatusPath + "?token=" + token
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var importStatusResp models.GlobalResponse
	err = json.Unmarshal(jsonResp, &importStatusResp)
	if err != nil {
		return nil, err
	}
	return &importStatusResp, nil
}
