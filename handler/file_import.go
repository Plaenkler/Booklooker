package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func ImportFile(token model.Token, req model.FileImportRequest) (*model.GlobalResponse, error) {
	url := model.BaseURL + model.FileImportPath + "?token=" + token.Value
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
	var fileImportResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &fileImportResp)
	if err != nil {
		return nil, err
	}
	return &fileImportResp, nil
}
