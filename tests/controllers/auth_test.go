package controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"smartjobsolutions/database"
	"smartjobsolutions/routes"
	"smartjobsolutions/types"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func testSetup() {
	database.InitDB()
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("testSetup error loading dotenv: ", err)
	}
}

func Test_SignUp(t *testing.T) {
	testSetup()
	userDetails := types.UserDetails{
		Username: "Raphael",
		Email:    "raphael@gmail.com",
		Password: "hashedPassword",
		Location: "Komarock",
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

func Test_SignIn(t *testing.T) {
	testSetup()
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
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_RegisterServiceProvider(t *testing.T) {
	testSetup()
	serviceProvider := types.ServiceProvider{
		Id:          uuid.New().String(),
		Service:     "photography",
		Description: "3d photographs",
	}
	jsonData, err := json.Marshal(serviceProvider)
	log.Print("json data: ", jsonData)
	if err != nil {
		t.Errorf("Test_RegisterServiceProvider Error: Could not parse json: %s", err)
	}
	req, err := http.NewRequest("POST", "/sign-up/service-provider", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Test_RegisterServiceProvider Error: could not send request: %s", err)
	}
	rr := httptest.NewRecorder()
	router := routes.SetupRouter()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
