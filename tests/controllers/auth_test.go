package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
