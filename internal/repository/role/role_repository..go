package repository

import (
	"database/sql"
	"todo-restful-api/internal/model/domain"
)

type RoleRepository interface {
	Create(db *sql.DB, role domain.Role) domain.Role
	Update(db *sql.DB, role domain.Role) domain.Role
	Delete(db *sql.DB, roleId int)
	FindById(db *sql.DB, roleId int) (domain.Role, error)
	FindAll(db *sql.DB) []domain.Role
}
