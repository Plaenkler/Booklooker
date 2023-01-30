package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/api/models"
)

func ImportFile(token string, req models.FileImportRequest) (*models.GeneralResponse, error) {
	url := baseURL + models.FileImportPath + "?token=" + token
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
	var fileImportResp models.GeneralResponse
	err = json.Unmarshal(jsonResp, &fileImportResp)
	if err != nil {
		return nil, err
	}
	return &fileImportResp, nil
}
