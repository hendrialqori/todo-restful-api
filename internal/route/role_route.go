package route

import (
	"database/sql"
	controller "todo-restful-api/internal/controller/role"
	repository "todo-restful-api/internal/repository/role"
	service "todo-restful-api/internal/service/role"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func RoleRouter(router *httprouter.Router, DB *sql.DB, validate *validator.Validate) {

	var (
		roleRepository = repository.NewRoleRepository(DB)
		roleService    = service.NewRoleService(roleRepository, validate)
		roleController = controller.NewRoleController(roleService)
	)

	router.GET("/api/roles", roleController.FindAll)
	router.GET("/api/roles/:roleId", roleController.FindById)
	router.POST("/api/roles", roleController.Create)
	router.PUT("/api/roles/:roleId", roleController.Update)
	router.DELETE("/api/roles/:roleId", roleController.Delete)
}
