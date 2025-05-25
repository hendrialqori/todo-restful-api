package service

import "todo-restful-api/internal/model/web"

type AuthService interface {
	Login(request web.AuthLoginRequest) web.AuthResponse
}
