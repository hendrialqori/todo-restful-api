package repository

import (
	"todo-restful-api/internal/model/domain"
)

type UserRepository interface {
	Create(user domain.User) domain.User
	Update(user domain.User) domain.User
	Delete(userId int)
	FindById(userId int) (domain.User, error)
	FindAll() []domain.User
}
