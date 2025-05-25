package helper

import (
	"fmt"
	"time"
	"todo-restful-api/internal/model/domain"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("OMKEGAS!")

func CreatToken(user domain.User) (string, error) {
	var (
		token string
		err   error
	)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.UserName,
		"email":    user.Email,
		"role":     user.Role,
		"iss":      "todo_restful_api",
		"sub":      "123456",
		"exp":      time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":      time.Now().Unix(),
	})

	if token, err = claims.SignedString(secretKey); err != nil {
		return token, err
	}
	return token, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	var (
		token *jwt.Token
		err   error
	)

	if token, err = jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// Pastikan algoritma sesuai ekspektasi
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	}); err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return token, nil
}
