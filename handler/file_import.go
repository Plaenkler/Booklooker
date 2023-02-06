package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/plaenkler/booklooker/model"
)

func ImportFile(token model.Token, req model.FileImportRequest) (*model.GlobalResponse, error) {
	if token.Value == "" {
		return nil, fmt.Errorf("token has no value")
	}
	if token.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	params := url.Values{}
	if req.File == nil {
		return nil, fmt.Errorf("file is required")
	}
	if req.FileType != "" {
		params.Set("fileType", string(req.FileType))
	}
	if req.MediaType != "" {
		params.Set("mediaType", string(req.MediaType))
	}
	if req.DataType != 1 {
		params.Set("dataType", fmt.Sprintf("%d", req.DataType))
		fmt.Println(params)
	}
	if req.FormatID != "" {
		params.Set("formatId", req.FormatID)
	}
	if req.Encoding != "" {
		params.Set("encoding", string(req.Encoding))
	}
	url := model.BaseURL + model.FileImportPath + "?token=" + token.Value + "&" + params.Encode()
	fmt.Println(url)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", req.File.Name())
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, req.File)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, writer.FormDataContentType(), body)
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
