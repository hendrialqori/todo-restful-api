package exception

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"todo-restful-api/internal/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

func PanicHandler(w http.ResponseWriter, r *http.Request, err any) {
	fmt.Println(err)

	if sqlError(w, err) {
	}

	if validationError(w, err) {
	}

	if notFoundError(w, err) {
	}
}

func sqlError(w http.ResponseWriter, err any) bool {
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
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, err any) bool {
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
		return true

	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, err any) bool {
	if _, ok := err.(NotFoundError); ok {

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

		return true
	} else {
		return false
	}
}
