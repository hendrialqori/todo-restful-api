package route

import (
	"database/sql"
	controller "todo-restful-api/internal/controller/user"
	repository "todo-restful-api/internal/repository/user"
	service "todo-restful-api/internal/service/user"
	"todo-restful-api/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func UserRouter(router *httprouter.Router, DB *sql.DB, validate *validator.Validate) {

	var (
		userRepository = repository.NewUserRepository(DB)
		userService    = service.NewUserService(userRepository, validate)
		userController = controller.NewUserController(userService)
	)

	router.GET("/api/users", middleware.Auth(userController.FindAll))
	router.GET("/api/users/:userId", middleware.Auth(userController.FindById))
	router.POST("/api/users", middleware.Auth(userController.Create))
	router.PUT("/api/users/:userId", middleware.Auth(userController.Update))
	router.DELETE("/api/users/:userId", middleware.Auth(userController.Delete))
}
