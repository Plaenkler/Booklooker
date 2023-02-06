package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/plaenkler/booklooker/model"
)

func GetFileStatus(token model.Token, req model.FileStatusRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	if req.Filename == "" {
		return nil, fmt.Errorf("filename is required")
	}
	url := model.BaseURL + model.FileStatusPath + "?token=" + token.Value + "&filename=" + req.Filename
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var fileStatusResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &fileStatusResp)
	if err != nil {
		return nil, err
	}
	return &fileStatusResp, nil
}
