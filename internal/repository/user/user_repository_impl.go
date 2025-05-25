package repository

import (
	"database/sql"
	"encoding/json"
	"todo-restful-api/internal/model/domain"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

// Create implements UserRepository.
func (repository *UserRepositoryImpl) Create(user domain.User) domain.User {
	query := `
		INSERT INTO users(email, username, role_id) VALUES(?, ?, ?)
	`
	result, err := repository.DB.Exec(query, user.Email, user.UserName, user.RoleId)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = int(id)

	return user
}

// Delete implements UserRepository.
func (repository *UserRepositoryImpl) Delete(userId int) {
	query := `delete from users where id = (?)`
	if _, err := repository.DB.Exec(query, userId); err != nil {
		panic(err)
	}
}

// FindAll implements UserRepository.
func (repository *UserRepositoryImpl) FindAll() []domain.User {
	query := `
		SELECT u.id, u.email, u.username, json_object('id',r.id,'name',r.name) AS ROLE,
		 u.created_at, u.updated_at
		FROM users AS u
		JOIN roles AS r ON (u.role_id = r.id)
	`
	rows, err := repository.DB.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		var roleJSON string

		if err := rows.Scan(&user.Id, &user.Email, &user.UserName,
			&roleJSON, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		if err := json.Unmarshal([]byte(roleJSON), &user.Role); err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

// FindById implements UserRepository.
func (repository *UserRepositoryImpl) FindById(userId int) (domain.User, error) {
	query := `
		SELECT u.id, u.email, u.username, json_object('id',r.id,'name',r.name) AS ROLE,
		u.created_at, u.updated_at
		FROM users AS u
		JOIN roles AS r ON (u.role_id = r.id)
		WHERE u.id = (?)
	`
	row := repository.DB.QueryRow(query, userId)

	var user domain.User
	var roleJSON string

	if err := row.Scan(&user.Id, &user.Email, &user.UserName,
		&roleJSON, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	if err := json.Unmarshal([]byte(roleJSON), &user.Role); err != nil {
		return user, err
	}

	return user, nil
}

// Update implements UserRepository.
func (repository *UserRepositoryImpl) Update(user domain.User) domain.User {
	query := `
		UPDATE users SET email = (?), username = (?), updated_at = current_timestamp WHERE id = (?)
	`
	if _, err := repository.DB.Exec(query, user.Email, user.UserName, user.Id); err != nil {
		panic(err)
	}

	return user

}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}
