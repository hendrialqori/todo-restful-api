package route

import (
	"database/sql"
	controller "todo-restful-api/internal/controller/todo"
	repository "todo-restful-api/internal/repository/todo"
	service "todo-restful-api/internal/service/todo"
	"todo-restful-api/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func TodoRouter(router *httprouter.Router, DB *sql.DB, validate *validator.Validate) {

	var (
		roleRepository = repository.NewTodoRepository(DB)
		roleService    = service.NewTodoService(roleRepository, validate)
		roleController = controller.NewTodoController(roleService)
	)

	router.GET("/api/todos/:userId", middleware.Auth(roleController.FindAll))
	router.POST("/api/todos", middleware.Auth(roleController.Create))
	router.PUT("/api/todos/:todoId", middleware.Auth(roleController.Update))
	router.DELETE("/api/todos/:todoId/users/:userId", middleware.Auth(roleController.Delete))
}
