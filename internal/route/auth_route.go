package route

import (
	"database/sql"
	controller "todo-restful-api/internal/controller/auth"
	repository "todo-restful-api/internal/repository/auth"
	service "todo-restful-api/internal/service/auth"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func AuthRouter(router *httprouter.Router, DB *sql.DB, validate *validator.Validate) {

	var (
		roleRepository = repository.NewAuthRepository(DB)
		roleService    = service.NewAuthService(roleRepository, validate)
		roleController = controller.NewAuthController(roleService)
	)

	router.POST("/api/auth/login", roleController.Login)
}
