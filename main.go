package main

import (
	"net/http"
	"todo-restful-api/app"
	"todo-restful-api/exception"
	"todo-restful-api/internal/route"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var (
		router   = httprouter.New()
		DB       = app.NewDB()
		validate = validator.New()
	)

	route.RoleRouter(router, DB, validate)
	route.UserRouter(router, DB, validate)

	router.PanicHandler = exception.PanicHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
