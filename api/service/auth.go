package service

import (
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"articleproject/api/repository"
)

type AuthService interface {
	UserRegistration(user request.User) (error) 
	UserLogin(user request.User) (response.User, string, error)
	RefreshToken(string) (int64, bool, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(a repository.AuthRepository) AuthService {
	return authService{
		authRepository: a,
	}
}

func (a authService) UserRegistration(user request.User) (error){
	return a.authRepository.UserRegistration(user)
}

func (a authService) UserLogin(user request.User) (response.User, string, error) {
	return a.authRepository.UserLogin(user)
}

func (a authService) RefreshToken(token string) (int64, bool, error) {
	return a.authRepository.RefreshToken(token)
}
