package repository

import (
	"todo-restful-api/internal/model/domain"
)

type TodoRepository interface {
	Create(todo domain.Todo) domain.Todo
	Update(todo domain.Todo) domain.Todo
	Delete(todoId int, userId int)
	FindById(todoId int) (domain.Todo, error)
	FindAll(todoId int) []domain.Todo
}
