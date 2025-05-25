package repository

import "todo-restful-api/internal/model/domain"

type AuthRepository interface {
	Login(user domain.User) domain.User
}
