package main

import (
	"net/http"
	"todo-restful-api/app"
	"todo-restful-api/exception"
	"todo-restful-api/internal/route"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"

	_ "todo-restful-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title						Todo Restful API
//	@version					1.0
//	@description				This is a sample todo api (CRUD).
//	@host						localhost:3000
//	@BasePath					/api
//	@SecurityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization

func main() {
	var (
		router   = httprouter.New()
		DB       = app.NewDB()
		validate = validator.New()
	)

	// Swagger documentation route
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	route.AuthRouter(router, DB, validate)
	route.RoleRouter(router, DB, validate)
	route.UserRouter(router, DB, validate)
	route.TodoRouter(router, DB, validate)

	router.PanicHandler = exception.PanicHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
