package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-restful-api/internal/model/web"
	service "todo-restful-api/internal/service/user"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

// Create implements UserController.
func (controller *UserControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// initiate struct user request
	var userRequest web.UserCreateRequest

	// parse request body with json NewDecoder method and assign to userRequest struct
	if err := json.NewDecoder(request.Body).Decode(&userRequest); err != nil {
		panic(err)
	}

	// passing userRequest into user service method
	userResponse := controller.UserService.Create(userRequest)

	// mapping userResponse into ApiResponse
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    userResponse,
	}

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// Delete implements UserController.
func (controller *UserControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}

	controller.UserService.Delete(id)

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

// FindAll implements UserController.
func (controller *UserControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersResponse := controller.UserService.FindAll()
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Sucess",
		Data:    usersResponse,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// FindById implements UserController.
func (controller *UserControllerImpl) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}

	userResponse := controller.UserService.FindById(id)
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userResponse,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// Update implements UserController.
func (controller *UserControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(err)
	}

	// initiate struct role request
	var userRequest web.UserUpdateRequest
	userRequest.Id = id

	// parse request body with json NewDecoder method and assign to userRequest struct
	if err = json.NewDecoder(request.Body).Decode(&userRequest); err != nil {
		panic(err)
	}

	// passing userRequest into role service method
	userResponse := controller.UserService.Update(userRequest)

	// mapping roleResponse into ApiResponse
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    userResponse,
	}

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

func NewUserController(UserService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: UserService,
	}
}
