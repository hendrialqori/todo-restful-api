package service

import (
	"todo-restful-api/exception"
	"todo-restful-api/internal/model/domain"
	"todo-restful-api/internal/model/web"
	repository "todo-restful-api/internal/repository/user"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

// Create implements UserService.
func (service *UserServiceImpl) Create(request web.UserCreateRequest) web.UserResponse {
	var userCreateRequest domain.User

	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	userCreateRequest = domain.User{Email: request.Email, UserName: request.UserName, RoleId: request.RoleId}
	userCreateRequest = service.UserRepository.Create(userCreateRequest)

	return web.UserResponse{
		Id:       userCreateRequest.Id,
		Email:    userCreateRequest.Email,
		UserName: userCreateRequest.UserName,
	}
}

// Delete implements UserService.
func (service *UserServiceImpl) Delete(UserId int) {
	if _, err := service.UserRepository.FindById(UserId); err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(UserId)
}

// FindAll implements UserService.
func (service *UserServiceImpl) FindAll() []web.UserResponse {
	users := service.UserRepository.FindAll()

	var usersResponse []web.UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, web.UserResponse{
			Id:       user.Id,
			Email:    user.Email,
			UserName: user.UserName,
			Role: web.RoleResponse{
				Id:   user.Role.Id,
				Name: user.Role.Name,
			},
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return usersResponse
}

// FindById implements UserService.
func (service *UserServiceImpl) FindById(UserId int) web.UserResponse {
	user, err := service.UserRepository.FindById(UserId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userResponse := web.UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		UserName: user.UserName,
		Role: web.RoleResponse{
			Id:   user.Role.Id,
			Name: user.Role.Name,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userResponse
}

// Update implements UserService.
func (service *UserServiceImpl) Update(request web.UserUpdateRequest) web.UserResponse {
	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	user, err := service.UserRepository.FindById(request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Email = request.Email
	user.UserName = request.UserName

	user = service.UserRepository.Update(user)
	userResponse := web.UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		UserName: user.UserName,
		Role: web.RoleResponse{
			Id:   user.Role.Id,
			Name: user.Role.Name,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse
}

func NewUserService(r repository.UserRepository, v *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: r,
		Validate:       v,
	}
}
