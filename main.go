package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"todo-restful-api/app"
	"todo-restful-api/exception"
	controller "todo-restful-api/internal/controller/role"
	"todo-restful-api/internal/model/web"
	repository "todo-restful-api/internal/repository/role"
	service "todo-restful-api/internal/service/role"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var (
		DB       = app.NewDB()
		router   = httprouter.New()
		validate = validator.New()

		roleRepository = repository.NewRoleRepository(DB)
		roleService    = service.NewRoleService(roleRepository, validate)
		roleController = controller.NewRoleController(roleService)
	)

	router.GET("/api/roles", roleController.FindAll)
	router.GET("/api/roles/:roleId", roleController.FindById)
	router.POST("/api/roles", roleController.Create)
	router.PUT("/api/roles/:roleId", roleController.Update)
	router.DELETE("/api/roles/:roleId", roleController.Delete)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err any) {
		fmt.Println(err)

		// Validator error
		if validatorError, ok := err.(validator.ValidationErrors); ok {
			w.Header().Add("X-Status-Reason", "Validation failed")
			w.WriteHeader(http.StatusUnprocessableEntity)

			errMessages := make([]string, 0)
			for _, fieldErr := range validatorError {
				msg := fmt.Sprintf("Field '%s' failed on '%s' validation", fieldErr.Field(), fieldErr.Tag())
				errMessages = append(errMessages, msg)
			}

			errResponse := web.ApiResponse{
				Ok:      false,
				Code:    http.StatusUnprocessableEntity,
				Message: strings.Join(errMessages, ", "),
			}

			if err := json.NewEncoder(w).Encode(errResponse); err != nil {
				panic(err)
			}
		}

		if sqlError, ok := err.(*mysql.MySQLError); ok {
			w.Header().Add("X-Status-Reason", "Entity failed")
			w.WriteHeader(http.StatusInternalServerError)
			errResponse := web.ApiResponse{
				Ok:      false,
				Code:    http.StatusInternalServerError,
				Message: sqlError.Message,
			}

			if err := json.NewEncoder(w).Encode(errResponse); err != nil {
				panic(err)
			}
		}

		if _, ok := err.(exception.NotFoundError); ok {
			w.Header().Add("X-Status-Reason", "Data not found")
			w.WriteHeader(http.StatusNotFound)
			errResponse := web.ApiResponse{
				Ok:      false,
				Code:    http.StatusNotFound,
				Message: "Not found error",
			}

			if err := json.NewEncoder(w).Encode(errResponse); err != nil {
				panic(err)
			}
		}
	}

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
