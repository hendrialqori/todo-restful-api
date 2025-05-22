package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-restful-api/internal/model/web"
	service "todo-restful-api/internal/service/todo"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

// Create implements TodoController.
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

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// Delete implements TodoController.
func (controller *TodoControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}

	controller.TodoService.Delete(id)

	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "success",
	}

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// FindAll implements TodoController.
func (controller *TodoControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}

	todosResponse := controller.TodoService.FindAll(id)

	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "success",
		Data:    todosResponse,
	}

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// Update implements TodoController.
func (controller *TodoControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var todoRequest web.TodoUpdateRequest

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(request.Body).Decode(&todoRequest); err != nil {
		panic(err)
	}

	todoRequest.Id = id

	todoResponse := controller.TodoService.Update(todoRequest)
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "success",
		Data:    todoResponse,
	}

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	if err := json.NewEncoder(write).Encode(apiResponse); err != nil {
		panic(err)
	}
}

func NewTodoController(TodoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: TodoService,
	}
}
