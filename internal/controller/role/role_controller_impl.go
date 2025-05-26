package controller

import (
	"encoding/json"
	"net/http"
	"todo-restful-api/helper"
	"todo-restful-api/internal/model/web"
	service "todo-restful-api/internal/service/role"

	"github.com/julienschmidt/httprouter"
)

type RoleControllerImpl struct {
	RoleService service.RoleService
}

// CreateRoles godoc
//
//	@Summary		Create new role
//	@Description	Create new role
//	@Tags			Roles
//	@Accept			json
//	@Produce		json
//	@Param			body	body		web.RoleCreateRequest	true	"all fields is require"
//	@Security 		Bearer
//	@Success		200		{object}	web.ApiResponse[web.RoleResponse]
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/roles [post]
func (controller *RoleControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var roleRequest web.RoleCreateRequest
	if err := json.NewDecoder(request.Body).Decode(&roleRequest); err != nil {
		panic(err)
	}

	roleResponse := controller.RoleService.Create(roleRequest)
	apiResponse := web.ApiResponse[web.RoleResponse]{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    roleResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// UpdatedRoles godoc
//
//	@Summary		Update role base on id
//	@Description	Update existing role base on id
//	@Tags			Roles
//	@Param			id		path	int						true	"role id"
//	@Param			body	body	web.RoleUpdateRequest	true	"ignore or delete id field on request body"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse[web.RoleResponse]
//	@Failure		404	{object}	web.NotFoundErrorResponse
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/roles/{id} [put]
func (controller *RoleControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := helper.ParamInt(params, "roleId")

	var roleRequest web.RoleUpdateRequest
	if err := json.NewDecoder(request.Body).Decode(&roleRequest); err != nil {
		panic(err)
	}

	roleRequest.Id = roleId
	roleResponse := controller.RoleService.Update(roleRequest)

	apiResponse := web.ApiResponse[web.RoleResponse]{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    roleResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// DeleteRolesById godoc
//
//	@Summary		Delete roles base on id
//	@Description	Delete existing roles base on id
//	@Tags			Roles
//	@Param			id	path	int	true	"role id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse[web.RoleResponse]
//	@Failure		404	{object}	web.NotFoundErrorResponse
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/roles/{id} [delete]
func (controller *RoleControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := helper.ParamInt(params, "roleId")

	controller.RoleService.Delete(roleId)
	apiResponse := web.ApiResponse[web.RoleResponse]{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Sucess",
	}
	helper.ResponseToJson(write, apiResponse)
}

// GetRolesById godoc
//
//	@Summary		Get role base on id
//	@Description	Retrieve existing role base on id
//	@Tags			Roles
//	@Param			id	path	int	true	"role id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse[web.RoleResponse]
//	@Failure		404	{object}	web.NotFoundErrorResponse
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/roles/{id} [get]
func (controller *RoleControllerImpl) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := helper.ParamInt(params, "roleId")

	roleResponse := controller.RoleService.FindById(roleId)
	apiResponse := web.ApiResponse[web.RoleResponse]{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    roleResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

// ListRoles godoc
//
//	@Summary		Get all role
//	@Description	Retrieve all existing role
//	@Tags			Roles
//	@Accept			json
//	@Produce		json
//	@Security 		Bearer
//	@Success		200	{object}	web.ApiResponse[[]web.RoleResponse]
//	@Failure		401	{object}	web.UnAuthorizedErrorResponse
//	@Failure		500	{object}	web.InternalServerErrorResponse
//	@Router			/roles [get]
func (controller *RoleControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rolesResponse := controller.RoleService.FindAll()
	apiResponse := web.ApiResponse[[]web.RoleResponse]{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Sucess",
		Data:    rolesResponse,
	}
	helper.ResponseToJson(write, apiResponse)
}

func NewRoleController(RoleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: RoleService,
	}
}
