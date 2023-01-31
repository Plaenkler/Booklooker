package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func GetFileStatus(token model.Token, req model.FileStatusRequest) (*model.GlobalResponse, error) {
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
