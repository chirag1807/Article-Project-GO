package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"articleproject/api/service"
	"articleproject/api/validation"
	"articleproject/constants"
	"articleproject/error"
	"articleproject/utils"
)

type AuthController interface {
	UserRegistration(w http.ResponseWriter, r *http.Request)
	UserLogin(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(s service.AuthService) AuthController {
	return authController{
		authService: s,
	}
}

func (a authController) UserRegistration(w http.ResponseWriter, r *http.Request) {
	var user request.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	// err := json.NewDecoder(r.Body).Decode(&user)

	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	isEmail := validation.EmailValidation(user.Email)
	if !isEmail {
		utils.ErrorGenerator(w, errorhandling.EmailvalidationError)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}
	user.Password = hashedPassword

	err = a.authService.UserRegistration(user)
	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.USER_REGISTRATION_SUCCEED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}

func (a authController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest request.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &userLoginRequest)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	isEmail := validation.EmailValidation(userLoginRequest.Email)
	if !isEmail {
		utils.ErrorGenerator(w, errorhandling.EmailvalidationError)
		return
	}

	var user response.User
	var refreshToken string
	user, refreshToken, err = a.authService.UserLogin(userLoginRequest)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.LoginFailedError)
		return
	}

	accessToken, err := utils.CreateAccessToken(time.Now().Add(time.Hour * 5), user.ID, user.IsAdmin)
	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.UserResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}

func (a authController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(string)

	id, isadmin, err := a.authService.RefreshToken(token)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.RefreshTokenError)
		return
	}

	accessToken, err := utils.CreateAccessToken(time.Now().Add(time.Hour * 5), id, isadmin)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.RefreshTokenError)
		return
	}

	response := response.AccessTokenResponse{
		AccessToken: accessToken,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}
