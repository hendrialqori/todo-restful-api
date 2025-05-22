package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	Create(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}
