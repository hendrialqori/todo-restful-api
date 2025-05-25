package route

import (
	"database/sql"
	controller "todo-restful-api/internal/controller/todo"
	repository "todo-restful-api/internal/repository/todo"
	service "todo-restful-api/internal/service/todo"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func TodoRouter(router *httprouter.Router, DB *sql.DB, validate *validator.Validate) {

	var (
		roleRepository = repository.NewTodoRepository(DB)
		roleService    = service.NewTodoService(roleRepository, validate)
		roleController = controller.NewTodoController(roleService)
	)

	router.GET("/api/todos/:userId", roleController.FindAll)
	router.POST("/api/todos", roleController.Create)
	router.PUT("/api/todos/:todoId", roleController.Update)
	router.DELETE("/api/todos/:todoId/users/:userId", roleController.Delete)
}
