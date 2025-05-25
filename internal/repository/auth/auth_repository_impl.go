package repository

import (
	"database/sql"
	"encoding/json"
	"todo-restful-api/internal/model/domain"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

// Login implements AuthRepository.
func (repository *AuthRepositoryImpl) Login(user domain.User) domain.User {
	query := `
		SELECT u.id, u.email, u.username, json_object('id',r.id,'name',r.name) AS ROLE,
		u.created_at, u.updated_at
		FROM users AS u
		JOIN roles AS r ON (u.role_id = r.id)
		WHERE u.email = (?) AND u.username = (?)
	`
	row := repository.DB.QueryRow(query, user.Email, user.UserName)

	var roleJSON string

	if err := row.Scan(&user.Id, &user.Email, &user.UserName,
		&roleJSON, &user.CreatedAt, &user.UpdatedAt); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(roleJSON), &user.Role); err != nil {
		panic(err)
	}

	return user
}

func NewAuthRepository(DB *sql.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: DB,
	}
}
