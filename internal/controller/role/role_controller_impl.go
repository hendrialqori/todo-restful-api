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

// Create implements RoleController.
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

// Update implements RoleController.
func (controller *RoleControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleId := params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	if err != nil {
		panic(err)
	}

	// initiate struct role request
	var roleRequest web.RoleUpdateRequest
	roleRequest.Id = id

	// parse request body with json NewDecoder method and assign to roleRequest struct
	err = json.NewDecoder(request.Body).Decode(&roleRequest)
	if err != nil {
		panic(err)
	}

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

// Delete implements RoleController.
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

// FindById implements RoleController.
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

// FindAll implements RoleController.
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
