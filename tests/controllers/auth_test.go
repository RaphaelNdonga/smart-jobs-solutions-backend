package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"smartjobsolutions/database"
	"smartjobsolutions/routes"
	"smartjobsolutions/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignUp(t *testing.T) {
	userDetails := types.UserDetails{
		Username: "Raphael",
		Email:    "raphael@gmail.com",
		Password: "hashedPassword",
		UserType: "Employee",
	}
	jsonData, err := json.Marshal(userDetails)
	if err != nil {
		t.Errorf("Could not connect to /sign-up: %s", err)
	}
	req, err := http.NewRequest("POST", "/sign-up", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Could not connect to /sign-up: %s", err)
	}

	rr := httptest.NewRecorder()
	database.InitDB()
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_SignIn(t *testing.T) {
	userDetails := types.UserDetails{
		Email:    "nkibi53@gmail.com",
		Password: "password",
	}
	jsonData, err := json.Marshal(userDetails)
	if err != nil {
		t.Errorf("Test_SignIn Error: Could not parse json: %s", err)
	}
	req, err := http.NewRequest("POST", "/sign-in", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Test_SignIn Error: could not send request: %s", err)
	}
	rr := httptest.NewRecorder()
	database.InitDB()
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
