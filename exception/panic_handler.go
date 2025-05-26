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

type ErrorHandlerFunc func(w http.ResponseWriter, err any) bool

var allErrors = []ErrorHandlerFunc{
	authError,
	sqlError,
	validationError,
	notFoundError,
}

func PanicHandler(w http.ResponseWriter, r *http.Request, err any) {
	fmt.Println(err)

	for _, handler := range allErrors {
		if errorHandler := handler(w, err); errorHandler {
			return
		}
	}
	// Fallback unknown error
	writeError(w, http.StatusInternalServerError, "Internal server error", web.InternalServerErrorResponse{
		Ok:      false,
		Code:    500,
		Message: "internal server error",
	})
}

func writeError(w http.ResponseWriter, status int, reason string, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-Status-Reason", reason)
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func authError(w http.ResponseWriter, err any) bool {
	if authError, ok := err.(AuthError); ok {
		writeError(w, http.StatusUnauthorized, "Unauthorized error", web.UnAuthorizedErrorResponse{
			Ok:      false,
			Code:    401,
			Message: authError.Error,
		})
		return true
	} else {
		return false
	}
}

func sqlError(w http.ResponseWriter, err any) bool {
	if sqlError, ok := err.(*mysql.MySQLError); ok {
		writeError(w, http.StatusInternalServerError, "Internal server error", web.InternalServerErrorResponse{
			Ok:      false,
			Code:    http.StatusInternalServerError,
			Message: sqlError.Error(),
		})
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, err any) bool {
	if validatorError, ok := err.(validator.ValidationErrors); ok {
		errMessages := make([]string, 0)
		for _, fieldErr := range validatorError {
			msg := fmt.Sprintf("Field '%s' failed on '%s' validation with error '%s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Error())
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
		writeError(w, http.StatusNotFound, "Not found error", web.NotFoundErrorResponse{
			Ok:      false,
			Code:    404,
			Message: notFoundError.Error,
		})
		return true
	} else {
		return false
	}
}
