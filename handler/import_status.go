package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/plaenkler/booklooker/model"
)

func GetImportStatus(token model.Token) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
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
