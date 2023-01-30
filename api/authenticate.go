package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://api.booklooker.de/2.0/authenticate"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func authenticate(apiKey string) (response, error) {
	var res response

	jsonData, err := json.Marshal(map[string]string{"apiKey": apiKey})
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return res, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
