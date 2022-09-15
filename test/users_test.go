package main

import (
	"encoding/json"
	"fmt"
	"gindemo/db"
	"gindemo/routes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	md "gindemo/middleware"

	"github.com/stretchr/testify/assert"
)

func init() {
	md.Init()
	db.Init()
}

func TestUsersCreateError(t *testing.T) {
	router := routes.SetRouter()

	w := httptest.NewRecorder()

	payload := strings.NewReader(`{
		"account": "Jack",
		"password": "12345"
	  }`)
	req, _ := http.NewRequest("POST", "/api/v1/users/create", payload)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["errMsg"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "account already exist", value)
}

func TestUsersCreateOk(t *testing.T) {
	router := routes.SetRouter()

	w := httptest.NewRecorder()
	payload := strings.NewReader(`{
		"account": "unit Testing",
		"password": "12345"
	  }`)
	req1, _ := http.NewRequest("POST", "/api/v1/users/create", payload)

	router.ServeHTTP(w, req1)
	//fmt.Println("w : " + w.Body.String())
	fmt.Println("w : " + strconv.Itoa(w.Code))
	assert.Equal(t, 200, w.Code)

	var response md.AuthToken
	// Convert the JSON response to a map
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	// Grab the value
	// Make some assertions on the correctness of the response.
	value1 := response.AccessExpAt
	assert.NotEqual(t, 0, value1)
	value2 := response.RefreshExpAt
	assert.NotEqual(t, 0, value2)
	value3 := response.AccessToken
	assert.NotEqual(t, "", value3)
	value4 := response.RefreshToken
	assert.NotEqual(t, "", value4)
	value5 := response.TokenType
	assert.NotEqual(t, "", value5)
}

func TestUsersLogin(t *testing.T) {

	router := routes.SetRouter()

	w := httptest.NewRecorder()
	//var jsonStr = []byte(`{"account":"Jack","password":"12345"}`)
	payload := strings.NewReader(`{
		"account": "Jack",
		"password": "12345"
	  }`)
	req, _ := http.NewRequest("POST", "/api/v1/users/login", payload)
	router.ServeHTTP(w, req)
	fmt.Println("Body : " + w.Body.String())
	assert.Equal(t, 200, w.Code)

	var response md.AuthToken
	// Convert the JSON response to a map
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	// Grab the value
	// Make some assertions on the correctness of the response.
	value1 := response.AccessExpAt
	assert.NotEqual(t, 0, value1)
	value2 := response.RefreshExpAt
	assert.NotEqual(t, 0, value2)
	value3 := response.AccessToken
	assert.NotEqual(t, "", value3)
	value4 := response.RefreshToken
	assert.NotEqual(t, "", value4)
	value5 := response.TokenType
	assert.NotEqual(t, "", value5)
}
