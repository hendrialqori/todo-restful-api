package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Login(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}
