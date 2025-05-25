package controller

import (
	"encoding/json"
	"net/http"
	"todo-restful-api/helper"
	"todo-restful-api/internal/model/web"
	service "todo-restful-api/internal/service/todo"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

// CreateTodos godoc
//
//	@Summary		Create new todo
//	@Description	Create new todo
//	@Tags			Todos
//	@Accept			json
//	@Produce		json
//	@Param			body	body	web.TodoCreateRequest	true	"all fields is require"
//	@Success		200	{object}	web.ApiResponse
//	@Router			/todos [post]
func (controller *TodoControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var todoRequest web.TodoCreateRequest
	if err := json.NewDecoder(request.Body).Decode(&todoRequest); err != nil {
		panic(err)
	}

	todoResponse := controller.TodoService.Create(todoRequest)
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "success",
		Data:    todoResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// DeleteTodosById godoc
//
//	@Summary		Delete todo base on id
//	@Description	Delete existing todo base on id and user id
//	@Tags			Todos
//	@Param			todo_id		path	int	true	"todo id"
//	@Param			user_id		path	int	true	"user id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse
//	@Failure		404	{object}	web.NotFoundResponse
//	@Router			/todos/{todo_id}/users/{user_id} [delete]
func (controller *TodoControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := helper.ParamInt(params, "todoId")
	userId := helper.ParamInt(params, "userId")
	controller.TodoService.Delete(todoId, userId)

	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "success",
	}
	helper.ResponseToJson(write, apiResponse)
}

// GetTodosById godoc
//
//	@Summary		Get all todo base on user id
//	@Description	Retrieve existing todo base on user id
//	@Tags			Todos
//	@Param			user_id	path	int	true	"user id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse
//	@Failure		404	{object}	web.NotFoundResponse
//	@Router			/todos/{user_id} [get]
func (controller *TodoControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.ParamInt(params, "userId")
	todosResponse := controller.TodoService.FindAll(userId)

	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "success",
		Data:    todosResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// UpdatedTodos godoc
//
//	@Summary		Update todo base on id
//	@Description	Update existing todo base on id
//	@Tags			Todos
//	@Param			role_id		path	int						true	"todo id"
//	@Param			user_id		path	int						true	"user id"
//	@Param			body	body	web.TodoUpdateRequest	true	"ignore or delete id field on request body"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse
//	@Failure		404	{object}	web.NotFoundResponse
//	@Router			/todos/{id} [put]
func (controller *TodoControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var todoRequest web.TodoUpdateRequest

	todoId := helper.ParamInt(params, "todoId")

	if err := json.NewDecoder(request.Body).Decode(&todoRequest); err != nil {
		panic(err)
	}
	todoRequest.Id = todoId

	todoResponse := controller.TodoService.Update(todoRequest)
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "success",
		Data:    todoResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

func NewTodoController(TodoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: TodoService,
	}
}
