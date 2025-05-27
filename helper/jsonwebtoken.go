package helper

import (
	"fmt"
	"strings"
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
		"role":     user.Role.Name,
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
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	}); err != nil {
		return nil, fmt.Errorf("could not JSON decode header")
	}

	if !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return token, nil
}

func GetTokenHeader(authHeader string) (string, error) {
	if len(authHeader) == 0 {
		return "", fmt.Errorf("authorization header needed")
	}

	var token string

	tokenSlice := strings.Split(authHeader, " ")

	if len(tokenSlice) == 0 {
		return "", fmt.Errorf("token needed")
	}

	token = tokenSlice[1]

	return token, nil
}
