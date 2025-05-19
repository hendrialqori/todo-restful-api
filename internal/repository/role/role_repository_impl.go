package repository

import (
	"database/sql"
	"errors"
	"todo-restful-api/internal/model/domain"
)

type RoleRepositoryImpl struct{}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

// Delete implements RoleRepository.
func (repository *RoleRepositoryImpl) Delete(db *sql.DB, roleId int) {
	panic("unimplemented")
}

// FindAll implements RoleRepository.
func (repository *RoleRepositoryImpl) FindAll(db *sql.DB) []domain.Role {
	panic("unimplemented")
}

// FindById implements RoleRepository.
func (repository *RoleRepositoryImpl) FindById(db *sql.DB, roleId int) (domain.Role, error) {
	query := "SELECT * FROM roles WHERE id = (?)"
	row := db.QueryRow(query, roleId)

	var role = domain.Role{}

	err := row.Scan(&role.Id, &role.Name)

	if err != nil {
		return role, nil
	} else {
		return role, errors.New("role not found")
	}

}

// Update implements RoleRepository.
func (repository *RoleRepositoryImpl) Update(db *sql.DB, role domain.Role) domain.Role {
	panic("unimplemented")
}

func (repository *RoleRepositoryImpl) Create(db *sql.DB, role domain.Role) domain.Role {
	query := "INSERT INTO ROLES(NAME) VALUES(?)"
	result, err := db.Exec(query, role.Name)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	role.Id = int(id)

	return role
}
