package repository

import (
	"todo-restful-api/internal/model/domain"
)

type RoleRepository interface {
	Create(role domain.Role) domain.Role
	Update(role domain.Role) domain.Role
	Delete(roleId int)
	FindById(roleId int) (domain.Role, error)
	FindAll() []domain.Role
}
