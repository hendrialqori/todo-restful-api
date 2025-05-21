package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"todo-restful-api/internal/model/domain"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

// Create implements UserRepository.
func (repository *UserRepositoryImpl) Create(user domain.User) domain.User {
	query := `
		insert into users(email, username, role_id) values(?, ?, ?)
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
		select u.id, u.email, u.username, json_object('id',r.id,'name',r.name) as role,
		 u.created_at, u.updated_at
		from users as u
		join roles as r on (u.role_id = r.id)
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
		select u.id, u.email, u.username, json_object('id',r.id,'name',r.name) as role,
		u.created_at, u.updated_at
		from users as u
		join roles as r on (u.role_id = r.id)
		where u.id = (?)
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
		update users set email = (?), username = (?), updated_at = current_timestamp where id = (?)
	`
	if _, err := repository.DB.Exec(query, user.Email, user.UserName, user.Id); err != nil {
		fmt.Println("Error here!")
		panic(err)
	}

	return user

}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}
