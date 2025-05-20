package service

import "todo-restful-api/internal/model/web"

type RoleService interface {
	Create(request web.RoleCreateRequest) web.RoleResponse
	Update(request web.RoleUpdateRequest) web.RoleResponse
	Delete(roleId int)
	FindById(roleId int) web.RoleResponse
	FindAll() []web.RoleResponse
}
