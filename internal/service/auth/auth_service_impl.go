package service

import (
	"todo-restful-api/helper"
	"todo-restful-api/internal/model/domain"
	"todo-restful-api/internal/model/web"
	repository "todo-restful-api/internal/repository/auth"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Validate       *validator.Validate
}

// Login implements AuthService.
func (service *AuthServiceImpl) Login(request web.AuthLoginRequest) web.AuthResponse {
	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	authLoginRequest := domain.User{
		Email:    request.Email,
		UserName: request.UserName,
	}
	authLoginRequest = service.AuthRepository.Login(authLoginRequest)

	token, err := helper.CreatToken(authLoginRequest)
	if err != nil {
		panic(err)
	}

	return web.AuthResponse{
		Token: token,
	}
}

func NewAuthService(r repository.AuthRepository, v *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: r,
		Validate:       v,
	}
}
