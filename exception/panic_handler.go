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

type ErrorHandler func(w http.ResponseWriter, err any) bool

var errorsType = []ErrorHandler{
	sqlError,
	validationError,
	notFoundError,
}

func PanicHandler(w http.ResponseWriter, r *http.Request, err any) {
	fmt.Println(err)

	for _, handler := range errorsType {
		if errorHandler := handler(w, err); errorHandler {
			return
		}
	}

	// Fallback unknown error
	writeError(w, http.StatusInternalServerError, "Internal server error", "Something went wrong")
}

func writeError(w http.ResponseWriter, status int, reason string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-Status-Reason", reason)
	w.WriteHeader(status)

	errResponse := web.ApiResponse{
		Ok:      false,
		Code:    status,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		panic(err)
	}
}

func sqlError(w http.ResponseWriter, err any) bool {
	if sqlError, ok := err.(*mysql.MySQLError); ok {
		writeError(w, http.StatusInternalServerError, "Entity error", sqlError.Message)
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, err any) bool {
	if validatorError, ok := err.(validator.ValidationErrors); ok {
		errMessages := make([]string, 0)
		for _, fieldErr := range validatorError {
			msg := fmt.Sprintf("Field '%s' failed on '%s' validation", fieldErr.Field(), fieldErr.Tag())
			errMessages = append(errMessages, msg)
		}
		writeError(w, http.StatusUnprocessableEntity, "Validation error", strings.Join(errMessages, ", "))
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, err any) bool {
	if notFoundError, ok := err.(NotFoundError); ok {
		writeError(w, http.StatusNotFound, "Not found error", notFoundError.Error)
		return true
	} else {
		return false
	}
}
