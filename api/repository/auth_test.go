package repository

import (
	"articleproject/api/model/request"
	errorhandling "articleproject/error"
	"fmt"
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
	}{
		{
			TestCaseName: "Success",
			Name:         "Chirag",
			Bio:          "Junior Software Engineer",
			Email:        "chiragmakwana1807@gmail.com",
			Password:     "Chirag123$",
			Image:        "",
			IsAdmin:      false,
			Expected:     nil,
		},
		{
			TestCaseName: "Duplicate Email",
			Name:         "Chirag",
			Bio:          "Junior Software Engineer",
			Email:        "chiragmakwana1807@gmail.com",
			Password:     "Chirag123$",
			Image:        "",
			IsAdmin:      false,
			Expected:     errorhandling.DuplicateEmailFound,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			user := request.User{
				Name:     v.Name,
				Bio:      v.Bio,
				Email:    v.Email,
				Password: v.Password,
				Image:    v.Image,
				IsAdmin:  v.IsAdmin,
			}

			err := NewAuthRepo(conn, rdb, amqpConn).UserRegistration(user)
			fmt.Println(err)

			assert.Equal(t, v.Expected, err)
		})
	}
}

func TestUserLogin(t *testing.T) {
	testCases := []struct {
		TestCaseName string
		Email        string
		Password     string
		Expected     interface{}
	}{
		{
			TestCaseName: "Success",
			Email:        "nirajdarji@gmail.com",
			Password:     "Niraj123$",
			Expected:     nil,
		},
		{
			TestCaseName: "No User",
			Email:        "nirajdarji",
			Password:     "Niraj123$",
			Expected:     errorhandling.NoUserFound,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			user := request.User{
				Email:    v.Email,
				Password: v.Password,
			}
			_, _, err := NewAuthRepo(conn, rdb, amqpConn).UserLogin(user)
			fmt.Println(err)

			assert.Equal(t, v.Expected, err)
		})
	}
}

func TestRefreshToken(t *testing.T) {
	testCases := []struct {
		TestCaseName string
		Token        string
		Expected     interface{}
	}{
		{
			TestCaseName: "Success",
			Token:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk2MzMyMTQsImlkIjoiOTQ2ODU3NTI3ODg4MjQ4ODMzIiwiaXNhZG1pbiI6ZmFsc2V9.H1v6PMFSz7UcxFPGVll_jB_pfnbUhS66XNKc5TN5_-o",
			Expected:     nil,
		},
		{
			TestCaseName: "Refresh Token Error",
			Token:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			Expected:     errorhandling.RefreshTokenNotFound,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			_, _, err := NewAuthRepo(conn, rdb, amqpConn).RefreshToken(v.Token)

			assert.Equal(t, v.Expected, err)
		})
	}
}
