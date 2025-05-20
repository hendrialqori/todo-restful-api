package repository

import (
	"database/sql"
	"todo-restful-api/internal/model/domain"
)

type RoleRepositoryImpl struct {
	DB *sql.DB
}

func NewRoleRepository(DB *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		DB: DB,
	}
}

func (repository *RoleRepositoryImpl) Create(role domain.Role) domain.Role {
	query := "INSERT INTO roles(name) VALUES(?)"
	result, err := repository.DB.Exec(query, role.Name)

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

// Update implements RoleRepository.
func (repository *RoleRepositoryImpl) Update(role domain.Role) domain.Role {
	_, err := repository.DB.Exec("UPDATE roles SET name = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		role.Name,
		role.Id)

	if err != nil {
		panic(err)
	}

	return role
}

// Delete implements RoleRepository.
func (repository *RoleRepositoryImpl) Delete(roleId int) {
	query := "DELETE FROM roles WHERE id = (?)"
	_, err := repository.DB.Exec(query, roleId)

	if err != nil {
		panic(err)
	}

}

// FindById implements RoleRepository.
func (repository *RoleRepositoryImpl) FindById(roleId int) (domain.Role, error) {
	query := "SELECT * FROM roles WHERE id = (?) LIMIT 1"
	row := repository.DB.QueryRow(query, roleId)

	var role domain.Role

	if err := row.Scan(&role.Id, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return role, err
	}

	return role, nil

}

// FindAll implements RoleRepository.
func (repository *RoleRepositoryImpl) FindAll() []domain.Role {
	query := "SELECT * FROM roles"
	rows, err := repository.DB.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var roles []domain.Role

	for rows.Next() {
		var role domain.Role

		if err := rows.Scan(&role.Id, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
			panic(err)
		}

		roles = append(roles, role)
	}

	return roles

}
