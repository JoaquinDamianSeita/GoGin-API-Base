package handlers

import (
	"GoGin-API-Base/dao"
	testhelpers "GoGin-API-Base/test_helpers"
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockUserService struct{}

func (m *MockUserService) RegisterUser(c *gin.Context) {
	var request dao.User
	c.ShouldBindJSON(&request)
	if request.Username == "" || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters."})
		return
	}

	log.Println(request.Username, request.Email)

	if request.Username == "invalid.user" || request.Email == "invalid.user@example.com" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the email or the user is already in use"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully created."})
}

func (m *MockUserService) LoginUser(c *gin.Context) {
	var request dao.User
	c.ShouldBindJSON(&request)
	if request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters."})
		return
	}

	if request.Email == "test.user@example.com" && request.Password == "password123" {
		c.JSON(http.StatusOK, gin.H{"token": "token", "expires_in": "3600"})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
}

func TestUserHandlerImpl_RegisterUser(t *testing.T) {
	userService := &MockUserService{}
	userHandler := UserHandlerInit(userService)
	serviceUri := "/api/users"

	var tests = []testhelpers.TestStructure{
		{
			Name:         "when the request is successful",
			Params:       `{"username": "test.user", "email": "test@example.com", "password": "password123"}`,
			ExpectedCode: http.StatusOK,
			ExpectedBody: `{"message":"User successfully created."}`,
		},
		{
			Name:         "when email is not present",
			Params:       `{"username": "test.user", "password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when email is empty",
			Params:       `{"email": "", "username": "test.user", "password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when password is not present",
			Params:       `{"email": "test.user@example.com", "username": "user.test"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when password is empty",
			Params:       `{"email": "test.user@example.com", "username": "test.user", "password": ""}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when username is not present",
			Params:       `{"email": "test.user@example.com", "password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when username is empty",
			Params:       `{"email": "test.user@example.com", "username": "", "password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when email is not available",
			Params:       `{"email": "invalid.user@example.com", "username": "test.user", "password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"the email or the user is already in use\"}",
		},
		{
			Name:         "when username is not available",
			Params:       `{"email": "test.user@example.com", "username": "invalid.user", "password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"the email or the user is already in use\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx, responseRecorder := testhelpers.MockPostRequest(tt.Params, serviceUri)

			userHandler.RegisterUser(ctx)

			testhelpers.AssertExpectedCodeAndBodyResponse(t, tt, responseRecorder)
		})
	}
}

func TestUserHandlerImpl_LoginUser(t *testing.T) {
	userService := &MockUserService{}
	userHandler := UserHandlerInit(userService)
	serviceUri := "/api/users/login"

	var tests = []testhelpers.TestStructure{
		{
			Name:         "when the request is successful",
			Params:       `{"email": "test.user@example.com", "password": "password123"}`,
			ExpectedCode: http.StatusOK,
			ExpectedBody: "{\"expires_in\":\"3600\",\"token\":\"token\"}",
		},
		{
			Name:         "when email is not present",
			Params:       `{"password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when email is empty",
			Params:       `{"email": "", "password": "password123"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when password is not present",
			Params:       `{"email": "test.user@example.com"}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "when password is empty",
			Params:       `{"email": "test.user@example.com", "password": ""}`,
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "{\"error\":\"Invalid parameters.\"}",
		},
		{
			Name:         "with invalid email",
			Params:       `{"email": "invalid.user@example.com", "password": "password123"}`,
			ExpectedCode: http.StatusUnauthorized,
			ExpectedBody: "{\"error\":\"invalid credentials\"}",
		},
		{
			Name:         "with invalid password",
			Params:       `{"email": "test.user@example.com", "password": "invalidpassword"}`,
			ExpectedCode: http.StatusUnauthorized,
			ExpectedBody: "{\"error\":\"invalid credentials\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx, responseRecorder := testhelpers.MockPostRequest(tt.Params, serviceUri)

			userHandler.LoginUser(ctx)

			testhelpers.AssertExpectedCodeAndBodyResponse(t, tt, responseRecorder)
		})
	}
}
