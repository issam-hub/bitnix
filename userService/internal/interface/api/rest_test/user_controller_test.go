package resttest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/application/command"
	"user-service/internal/application/common"
	"user-service/internal/interface/api/rest"
	"user-service/internal/tests/mocks"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	e := echo.New()

	t.Run("happy case", func(t *testing.T) {
		mockSvc := new(mocks.MockAccountsService)

		reqBody := map[string]any{
			"username": "test_test",
			"email":    "test@example.com",
			"password": "password123",
			"role":     "developer",
		}

		reqBodyBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/signup", bytes.NewReader(reqBodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewUserController(e, mockSvc)

		userID := uuid.New()

		token := "dGVzdF90ZXN0OnBhc3N3b3JkMTIz"

		registerUserCommandResult := &command.RegisterUserCommandResult{
			Result: &common.UserResult{
				ID:       userID,
				Username: "test_test",
				Email:    "test@example.com",
				Token:    token,
			},
		}

		mockSvc.On("Register", mock.Anything, mock.AnythingOfType("*command.RegisterUserCommand")).Return(registerUserCommandResult, nil)

		err := ctrl.RegisterUserController(c)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rec.Code)

		expectedResponseBody := map[string]any{
			"id":       userID.String(),
			"username": reqBody["username"],
			"email":    reqBody["email"],
			"token":    token,
		}

		var actualResponseBody map[string]any

		err = json.Unmarshal(rec.Body.Bytes(), &actualResponseBody)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		assert.Equal(t, expectedResponseBody, actualResponseBody)

		mockSvc.AssertExpectations(t)
	})

	t.Run("sad case", func(t *testing.T) {
		mockSvc := new(mocks.MockAccountsService)

		reqBody := map[string]any{
			"username": "test_test",
			"email":    "test@",
			"password": "password123",
			"role":     "developer",
		}

		reqBodyBytes, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/signup", bytes.NewReader(reqBodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		ctrl := rest.NewUserController(e, mockSvc)

		err := ctrl.RegisterUserController(c)

		assert.Error(t, err)

		httpErr, _ := err.(*echo.HTTPError)

		assert.Equal(t, http.StatusBadRequest, httpErr.Code)

		expectedError := map[string]string{
			"email": "email is invalid",
		}

		assert.Equal(t, expectedError, httpErr.Message)
	})
}

// func TestLoginUser(t *testing.T) {
// 	e := echo.New()

// 	t.Run("happy case", func(t *testing.T) {
// 		mockSvc := new(mocks.MockAccountsService)

// 		reqBody := map[string]any{
// 			"username": "test_test",
// 			"password": "password123",
// 		}

// 		reqBodyBytes, _ := json.Marshal(reqBody)

// 		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", bytes.NewReader(reqBodyBytes))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		rec := httptest.NewRecorder()

// 		c := e.NewContext(req, rec)

// 		ctrl := rest.NewUserController(e, mockSvc)

// 		userID := uuid.New()

// 		token := "dGVzdF90ZXN0OnBhc3N3b3JkMTIz"

// 		loginUserCommandResult := &command.LoginUserCommandResult{
// 			Result: &common.LoginUserResult{
// 				ID:    userID,
// 				Token: token,
// 			},
// 		}

// 		mockSvc.On("Login", mock.Anything, mock.AnythingOfType("*command.LoginUserCommand")).Return(loginUserCommandResult, nil)

// 		err := ctrl.LoginUserController(c)

// 		assert.NoError(t, err)

// 		assert.Equal(t, http.StatusCreated, rec.Code)

// 		expectedResponseBody := map[string]any{
// 			"id":    userID.String(),
// 			"token": token,
// 		}

// 		var actualResponseBody map[string]any

// 		err = json.Unmarshal(rec.Body.Bytes(), &actualResponseBody)
// 		if err != nil {
// 			t.Fatalf("failed to decode response body: %v", err)
// 		}

// 		assert.Equal(t, expectedResponseBody, actualResponseBody)

// 		mockSvc.AssertExpectations(t)
// 	})
// }
