package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
	// header = request.Header.Get("Authorization")
	// token  *string = &strings.Split(header, " ")[1]
	)

	// fmt.Println(token)
	middleware.Handler.ServeHTTP(writer, request)

	// jwt decode here
}
