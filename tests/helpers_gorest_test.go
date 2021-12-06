package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Weeping-Willow/bdd-go-testing/src"
	"github.com/stretchr/testify/assert"

	. "github.com/onsi/ginkgo"
)

var testConfig *src.Config

func contextName(mehtod, path string) string {
	return fmt.Sprintf("route being used %s %s", mehtod, path)
}

func cleanUpTestUser(email string) {
	if email == "" {
		email = getDefaultUser().Email
	}

	existingUser, status := getUser(email)
	if status != http.StatusNotFound && len(existingUser.Users) > 0 {
		deleteUser(existingUser.Users[0].Id)
	}
}

func deleteUser(id int) (string, int) {
	resp, err := src.SendRequest(http.MethodDelete, fmt.Sprintf(urlToTest+userDeleteRoute, id), testConfig.ApiToken, "")
	if err != nil {
		assert.NoError(GinkgoT(), err, "Error while sending the request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(GinkgoT(), err, "Error while parsing body")

	return string(body), resp.StatusCode
}

func updateUser(id int, u *user) (*user, string, int) {
	resp, err := src.SendRequest(http.MethodPatch, fmt.Sprintf(urlToTest+userPutRoute, id), testConfig.ApiToken, u)
	if err != nil {
		assert.NoError(GinkgoT(), err, "Error while sending the request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(GinkgoT(), err, "Error while parsing body")

	var userResponse *userCreateResponse
	json.Unmarshal(body, &userResponse)

	return &userResponse.User, string(body), resp.StatusCode
}

func getUser(email string) (*userListResponse, int) {
	resp, err := src.SendRequest(http.MethodGet, urlToTest+userListGetRoute, testConfig.ApiToken, getUserRequest{
		Email: email,
	})
	if err != nil {
		assert.NoError(GinkgoT(), err, "Error while sending the request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(GinkgoT(), err, "Error while parsing body")

	var userResponse *userListResponse
	json.Unmarshal(body, &userResponse)

	return userResponse, resp.StatusCode
}

func createUser(u *user) (*user, string, int) {
	resp, err := src.SendRequest(http.MethodPost, urlToTest+userPostRoute, testConfig.ApiToken, u)
	if err != nil {
		assert.NoError(GinkgoT(), err, "Error while sending the request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(GinkgoT(), err, "Error while parsing body")

	var userResponse *userCreateResponse
	json.Unmarshal(body, &userResponse)

	return &userResponse.User, string(body), resp.StatusCode

}
