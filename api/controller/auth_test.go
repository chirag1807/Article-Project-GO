package controller

import (
	testcase "articleproject/api/model/dto/test-cases"
	"articleproject/api/model/request"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRegistration(t *testing.T) {

	testCases := []struct {
		TestCaseName string
		Name         string
		Bio          string
		Email        string
		Password     string
		Image        string
		IsAdmin      bool
		Expected     interface{}
		StatusCode   int
	}{
		{
			TestCaseName: "Success",
			Name:         "Chirag",
			Bio:          "Junior Software Engineer",
			Email:        "chiragmakwana1807@gmail.com",
			Password:     "Chirag123$",
			Image:        "",
			IsAdmin:      false,
			Expected:     "User Registration Done Successfully.",
			StatusCode:   200,
		},
		{
			TestCaseName: "Invalid Email",
			Name:         "Chirag",
			Bio:          "Junior Software Engineer",
			Email:        "chiragmakwana1807",
			Password:     "Chirag123$",
			Image:        "",
			IsAdmin:      false,
			Expected:     "Email Validation Failed, Please Provide Valid Email.",
			StatusCode:   400,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			r.Post("/api/auth/registration", NewAuthController(authService).UserRegistration)

			user := request.User{
				Name:     v.Name,
				Bio:      v.Bio,
				Email:    v.Email,
				Password: v.Password,
				Image:    v.Image,
				IsAdmin:  v.IsAdmin,
			}
			jsonValue, _ := json.Marshal(user)
			req, _ := http.NewRequest("POST", "/api/auth/registration", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var response testcase.UserRegistration
			json.Unmarshal(w.Body.Bytes(), &response)
			// responseData, _ := io.ReadAll(w.Body)
			// str := string(responseData)

			assert.Equal(t, v.Expected, response.Message)
			assert.Equal(t, v.StatusCode, w.Code)
		})
	}
}

func TestUserLogin(t *testing.T) {
	testCases := []struct {
		TestCaseName string
		Email        string
		Password     string
		Expected     interface{}
		StatusCode   int
	}{
		{
			TestCaseName: "Success",
			Email:        "nirajdarji@gmail.com",
			Password:     "Niraj123$",
			Expected:     "User Login Done Successfully.",
			StatusCode:   200,
		},
		{
			TestCaseName: "Invalid Email",
			Email:        "nirajdarji",
			Password:     "Niraj123$",
			Expected:     "Email Validation Failed, Please Provide Valid Email.",
			StatusCode:   400,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			r.Post("/api/auth/login", NewAuthController(authService).UserLogin)

			user := request.User{
				Email:    v.Email,
				Password: v.Password,
			}
			jsonValue, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, v.StatusCode, w.Code)
		})
	}
}

func TestRefreshToken(t *testing.T) {
	testCases := []struct {
		TestCaseName string
		Token string
		StatusCode   int
	}{
		{
			TestCaseName: "Success",
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk2MzMyMTQsImlkIjoiOTQ2ODU3NTI3ODg4MjQ4ODMzIiwiaXNhZG1pbiI6ZmFsc2V9.H1v6PMFSz7UcxFPGVll_jB_pfnbUhS66XNKc5TN5_-o",
			StatusCode:   200,
		},
		{
			TestCaseName: "Refresh Token Error",
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			StatusCode:   401,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			r.Post("/api/auth/refresh-token", NewAuthController(authService).RefreshToken)

			req, _ := http.NewRequest("POST", "/api/auth/refresh-token", http.NoBody)
			ctx := context.WithValue(req.Context(), ContextKeyToken, v.Token)
			req = req.WithContext(ctx)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, v.StatusCode, w.Code)
		})
	}
}
