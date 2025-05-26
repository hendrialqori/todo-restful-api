package controller

import (
	"encoding/json"
	"net/http"
	"todo-restful-api/helper"
	"todo-restful-api/internal/model/web"
	service "todo-restful-api/internal/service/user"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

// CreateUsers godoc
//
//	@Summary		Create new user
//	@Description	Create new user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		web.UserCreateRequest	true	"all fields is require | username field shoud not contains whitespace"
//	@Success		200		{object}	web.ApiResponse[web.UserResponse]
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/users [post]
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
	apiResponse := web.ApiResponse[web.UserResponse]{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    userResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// DeleteUser godoc
//
//	@Summary		Delete user base on id
//	@Description	Delete existing user base on id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"user id"
//	@Success		200	{object}	web.ApiResponse[web.UserResponse]
//	@Failure		404	{object}	web.NotFoundErrorResponse
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/users/{id} [delete]
func (controller *UserControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.ParamInt(params, "userId")

	controller.UserService.Delete(userId)
	apiResponse := web.ApiResponse[web.UserResponse]{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "success",
	}
	helper.ResponseToJson(write, apiResponse)
}

// ListUsers godoc
//
//	@Summary		Get all user
//	@Description	Retrieve all existing user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse[[]web.UserResponse]
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/users [get]
func (controller *UserControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersResponse := controller.UserService.FindAll()
	apiResponse := web.ApiResponse[[]web.UserResponse]{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Sucess",
		Data:    usersResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// GetUserById godoc
//
//	@Summary		Get user base on id
//	@Description	Retrieve existing user base on id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"user id"
//	@Success		200	{object}	web.ApiResponse[web.UserResponse]
//	@Failure		404	{object}	web.NotFoundErrorResponse
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/users/{id} [get]
func (controller *UserControllerImpl) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.ParamInt(params, "userId")

	userResponse := controller.UserService.FindById(userId)
	apiResponse := web.ApiResponse[web.UserResponse]{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// UpdateUserById godoc
//
//	@Summary		Update user base on id
//	@Description	Update existing user base on id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"user id"
//	@Param			body	body		web.UserUpdateRequest	true	"ignore/delete id field on request body"
//	@Success		200		{object}	web.ApiResponse[web.UserResponse]
//	@Failure		404		{object}	web.NotFoundErrorResponse
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/users/{id} [put]
func (controller *UserControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := helper.ParamInt(params, "userId")

	// initiate struct role request
	var userRequest web.UserUpdateRequest
	userRequest.Id = userId

	// parse request body with json NewDecoder method and assign to userRequest struct
	if err := json.NewDecoder(request.Body).Decode(&userRequest); err != nil {
		panic(err)
	}

	// passing userRequest into role service method
	userResponse := controller.UserService.Update(userRequest)
	// mapping roleResponse into ApiResponse
	apiResponse := web.ApiResponse[web.UserResponse]{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    userResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

func NewUserController(UserService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: UserService,
	}
}
