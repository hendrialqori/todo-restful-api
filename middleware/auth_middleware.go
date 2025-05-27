package middleware

import (
	"net/http"
	"todo-restful-api/exception"
	"todo-restful-api/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

func Auth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		autorization := r.Header.Get("Authorization")
		tokenHeader, err := helper.GetTokenHeader(autorization)

		if err != nil {
			panic(exception.AuthError{Error: err.Error()})
		}

		tokenVerified, err := helper.VerifyToken(tokenHeader)

		_, ok := tokenVerified.Claims.(jwt.MapClaims)

		if !ok {
			panic(exception.AuthError{Error: "Cannot claim token"})
		}

		if err != nil {
			panic(exception.AuthError{Error: err.Error()})
		}

		h(w, r, ps)
	}
}
