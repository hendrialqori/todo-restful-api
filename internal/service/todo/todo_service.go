package service

import "todo-restful-api/internal/model/web"

type TodoService interface {
	Create(request web.TodoCreateRequest) web.TodoResponse
	Update(request web.TodoUpdateRequest) web.TodoResponse
	Delete(todoId int)
	FindAll(todoId int) []web.TodoResponse
}
