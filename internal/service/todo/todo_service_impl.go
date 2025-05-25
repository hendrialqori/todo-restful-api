package service

import (
	"todo-restful-api/constant"
	"todo-restful-api/exception"
	"todo-restful-api/internal/model/domain"
	"todo-restful-api/internal/model/web"
	repository "todo-restful-api/internal/repository/todo"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	Validate       *validator.Validate
}

// Create implements TodoService.
func (service *TodoServiceImpl) Create(request web.TodoCreateRequest) web.TodoResponse {
	var todoCreateRequest domain.Todo

	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	todoCreateRequest = domain.Todo{
		Title:  request.Title,
		UserId: request.UserId,
	}

	todoCreateRequest = service.TodoRepository.Create(todoCreateRequest)

	return web.TodoResponse{
		Id:     todoCreateRequest.Id,
		Title:  todoCreateRequest.Title,
		Status: constant.StatusOnPending,
	}
}

// Delete implements TodoService.
func (service *TodoServiceImpl) Delete(todoId int, userId int) {
	if _, err := service.TodoRepository.FindById(todoId); err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.TodoRepository.Delete(todoId, userId)
}

// FindAll implements TodoService.
func (service *TodoServiceImpl) FindAll(todoId int) []web.TodoResponse {
	todos := service.TodoRepository.FindAll(todoId)

	var todosResponse []web.TodoResponse

	for _, todo := range todos {
		todo := web.TodoResponse{
			Id:     todo.Id,
			Title:  todo.Title,
			Status: todo.Status,
			User: web.UserResponse{
				Id:       todo.User.Id,
				Email:    todo.User.Email,
				UserName: todo.User.UserName,
			},
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		}
		todosResponse = append(todosResponse, todo)
	}
	return todosResponse
}

// Update implements TodoService.
func (service *TodoServiceImpl) Update(request web.TodoUpdateRequest) web.TodoResponse {
	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	todo, err := service.TodoRepository.FindById(request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todo.Title = request.Title
	todo.Status = request.Status
	todo.UserId = request.UserId

	todo = service.TodoRepository.Update(todo)
	todoResponse := web.TodoResponse{
		Id:        todo.Id,
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return todoResponse
}

func NewTodoService(r repository.TodoRepository, v *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: r,
		Validate:       v,
	}
}
