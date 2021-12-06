package main

import (
	"encoding/json"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

func SendRequest(method, url, token string, body interface{}) *http.Response {
	var jsonStr []byte
	if body != nil {
		jsonStr, _ = json.Marshal(body)
	}

	req, err := http.NewRequest(method, url, strings.NewReader(string(jsonStr)))
	assert.NoError(GinkgoT(), err, "Error while creating request")

	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(GinkgoT(), err, "Error while sending request")

	return resp
}
