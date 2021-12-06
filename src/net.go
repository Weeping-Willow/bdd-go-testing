package src

import (
	"encoding/json"
	"net/http"
	"strings"
)

func SendRequest(method, url, token string, body interface{}) (*http.Response, error) {
	var jsonStr []byte
	if body != nil {
		jsonStr, _ = json.Marshal(body)
	}

	req, err := http.NewRequest(method, url, strings.NewReader(string(jsonStr)))
	if err != nil {
		return nil, err
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
