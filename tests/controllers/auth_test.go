package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"smartjobsolutions/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignUp(t *testing.T) {
	req, err := http.NewRequest("POST", "/sign-up", nil)
	if err != nil {
		t.Errorf("Could not connect to /sign-up: %s", err)
	}

	rr := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
