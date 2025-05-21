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
	query := `insert into roles(name) values(?)`
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
	query := `update roles set name = (?), updated_at = current_timestamp where id = (?)`
	if _, err := repository.DB.Exec(query, role.Name, role.Id); err != nil {
		panic(err)
	}
	return role
}

// Delete implements RoleRepository.
func (repository *RoleRepositoryImpl) Delete(roleId int) {
	query := `delete from roles where id = (?)`
	if _, err := repository.DB.Exec(query, roleId); err != nil {
		panic(err)
	}
}

// FindById implements RoleRepository.
func (repository *RoleRepositoryImpl) FindById(roleId int) (domain.Role, error) {
	query := `select * from roles where id = (?) limit 1`
	row := repository.DB.QueryRow(query, roleId)

	var role domain.Role

	if err := row.Scan(&role.Id, &role.Name, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return role, err
	}

	return role, nil
}

// FindAll implements RoleRepository.
func (repository *RoleRepositoryImpl) FindAll() []domain.Role {
	query := `select * from roles`
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
