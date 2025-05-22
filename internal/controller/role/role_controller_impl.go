package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-restful-api/internal/model/web"
	service "todo-restful-api/internal/service/role"

	"github.com/julienschmidt/httprouter"
)

type RoleControllerImpl struct {
	RoleService service.RoleService
}

// CreateRoles godoc
//
//	@Summary		Create roles
//	@Description	Create new role
//	@Tags			Roles
//	@Param			body	body	web.RoleCreateRequest	true "Role name must be capitalize"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse{}
//	@Failure		404	{object}	web.NotFoundResponse{}
//	@Router			/roles/{id} [post]
func (controller *RoleControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// initiate struct role request
	var roleRequest web.RoleCreateRequest

	// parse request body with json NewDecoder method and assign to roleRequest struct
	if err := json.NewDecoder(request.Body).Decode(&roleRequest); err != nil {
		panic(err)
	}

	// passing roleRequest into role service method
	roleResponse := controller.RoleService.Create(roleRequest)

	// mapping roleResponse into ApiResponse
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    roleResponse,
	}

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// UpdatedRoles godoc
//
//	@Summary		Update roles base on id
//	@Description	Update existing roles base on id
//	@Tags			Roles
//	@Param			id		path	int						true	"Role Id"
//	@Param			body	body	web.RoleUpdateRequest	true	"Data role yang akan diupdate (tanpa ID)"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse{}
//	@Failure		404	{object}	web.NotFoundResponse{}
//	@Router			/roles/{id} [put]
func (controller *RoleControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	if err != nil {
		panic(err)
	}

	// initiate struct role request
	var roleRequest web.RoleUpdateRequest

	// parse request body with json NewDecoder method and assign to roleRequest struct
	err = json.NewDecoder(request.Body).Decode(&roleRequest)
	if err != nil {
		panic(err)
	}

	roleRequest.Id = id

	// passing roleRequest into role service method
	roleResponse := controller.RoleService.Update(roleRequest)

	// mapping roleResponse into ApiResponse
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    roleResponse,
	}

	// add content type header response
	write.Header().Add("Content-Type", "application/json")

	// create encoder to encode apiResponse variable, then return into http response
	encode := json.NewEncoder(write)
	if err = encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// DeleteRolesById godoc
//
//	@Summary		Delete roles base on id
//	@Description	Delete existing roles base on id
//	@Tags			Roles
//	@Param			id	path	int	true	"Role Id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse{}
//	@Router			/roles/{id} [delete]
func (controller *RoleControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	if err != nil {
		panic(err)
	}

	controller.RoleService.Delete(id)

	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Sucess",
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}
}

// GetRolesById godoc
//
//	@Summary		Get roles base on id
//	@Description	Retrieve existing roles base on id
//	@Tags			Roles
//	@Param			id	path	int	true	"Role Id"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse{}
//	@Failure		404	{object}	web.NotFoundResponse{}
//	@Router			/roles/{id} [get]
func (controller *RoleControllerImpl) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	if err != nil {
		panic(err)
	}

	roleResponse := controller.RoleService.FindById(id)
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    roleResponse,
	}

	write.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(write).Encode(apiResponse); err != nil {
		panic(err)
	}
}

// ListRoles godoc
//
//	@Summary		Get roles
//	@Description	Retrieve all existing roles
//	@Tags			Roles
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.ApiResponse{}
//	@Failure		404	{object}	web.NotFoundResponse{}
//	@Router			/roles [get]
func (controller *RoleControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rolesResponse := controller.RoleService.FindAll()
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusOK,
		Message: "Sucess",
		Data:    rolesResponse,
	}

	write.Header().Add("Content-Type", "application/json")
	encode := json.NewEncoder(write)
	if err := encode.Encode(apiResponse); err != nil {
		panic(err)
	}

}

func NewRoleController(RoleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: RoleService,
	}
}
