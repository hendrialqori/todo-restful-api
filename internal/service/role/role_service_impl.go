package service

import (
	"todo-restful-api/exception"
	"todo-restful-api/internal/model/domain"
	"todo-restful-api/internal/model/web"
	repository "todo-restful-api/internal/repository/role"

	"github.com/go-playground/validator/v10"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Validate       *validator.Validate
}

// Create implements RoleService.
func (service *RoleServiceImpl) Create(request web.RoleCreateRequest) web.RoleResponse {
	var roleCreateRequest domain.Role

	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	roleCreateRequest = domain.Role{Name: request.Name}
	roleCreateRequest = service.RoleRepository.Create(roleCreateRequest)
	roleCreateRequest, _ = service.RoleRepository.FindById(roleCreateRequest.Id)

	return web.RoleResponse{
		Id:        roleCreateRequest.Id,
		Name:      roleCreateRequest.Name,
		CreatedAt: roleCreateRequest.CreatedAt,
		UpdatedAt: roleCreateRequest.UpdatedAt,
	}
}

// Delete implements RoleService.
func (service *RoleServiceImpl) Delete(roleId int) {
	_, err := service.RoleRepository.FindById(roleId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.RoleRepository.Delete(roleId)
}

// FindAll implements RoleService.
func (service *RoleServiceImpl) FindAll() []web.RoleResponse {
	roles := service.RoleRepository.FindAll()

	var rolesResponse []web.RoleResponse

	for _, role := range roles {
		rolesResponse = append(rolesResponse, web.RoleResponse{
			Id:        role.Id,
			Name:      role.Name,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
		})
	}

	return rolesResponse
}

// FindById implements RoleService.
func (service *RoleServiceImpl) FindById(roleId int) web.RoleResponse {
	role, err := service.RoleRepository.FindById(roleId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	roleResponse := web.RoleResponse{
		Id:        role.Id,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}

	return roleResponse
}

// Update implements RoleService.
func (service *RoleServiceImpl) Update(request web.RoleUpdateRequest) web.RoleResponse {
	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	role, err := service.RoleRepository.FindById(request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	role.Name = request.Name

	role = service.RoleRepository.Update(role)
	return web.RoleResponse{
		Id:        role.Id,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func NewRoleService(repository repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: repository,
		Validate:       validate,
	}
}
