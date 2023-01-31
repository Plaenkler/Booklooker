package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/plaenkler/booklooker/model"
)

func GetImportStatus(token model.Token) (*model.GlobalResponse, error) {
	url := model.BaseURL + model.ImportStatusPath + "?token=" + token.Value
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var importStatusResp model.GlobalResponse
	err = json.Unmarshal(jsonResp, &importStatusResp)
	if err != nil {
		return nil, err
	}
	return &importStatusResp, nil
}
