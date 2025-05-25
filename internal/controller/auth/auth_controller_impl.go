package controller

import (
	"encoding/json"
	"net/http"
	"todo-restful-api/helper"
	"todo-restful-api/internal/model/web"
	service "todo-restful-api/internal/service/auth"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

// AuthLogin godoc
//
//	@Summary		Auth login
//	@Description	Authentication with login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body	web.AuthLoginRequest	true	"all fields is require"
//	@Success		200	{object}	web.ApiResponse
//	@Router			/auth/login [post]
func (controller *AuthControllerImpl) Login(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var authLoginRequest web.AuthLoginRequest

	if err := json.NewDecoder(request.Body).Decode(&authLoginRequest); err != nil {
		panic(err)
	}

	authResponse := controller.AuthService.Login(authLoginRequest)
	apiResponse := web.ApiResponse{
		Ok:      true,
		Code:    http.StatusCreated,
		Message: "success",
		Data:    authResponse,
	}

	helper.ResponseToJson(write, apiResponse)
}

func NewAuthController(AuthService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: AuthService,
	}
}
