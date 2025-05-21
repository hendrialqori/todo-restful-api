package service

import "todo-restful-api/internal/model/web"

type UserService interface {
	Create(request web.UserCreateRequest) web.UserResponse
	Update(request web.UserUpdateRequest) web.UserResponse
	Delete(UserId int)
	FindById(UserId int) web.UserResponse
	FindAll() []web.UserResponse
}
