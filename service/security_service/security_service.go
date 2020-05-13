package security_service

import (
	"common"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"model/service_error"
	"net/http"
	"strings"
)

type config struct {
	Auth authConfig `yaml:"auth"`
}

type authConfig struct {
	Secret string `yaml:"secret"`
}

var secret string

func init() {
	cfg := config{}
	common.ReadConfig(&cfg)

	secret = cfg.Auth.Secret
}

func extractToken(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(authorization, "Bearer ")
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(secret), nil
}

func VerifyToken(r *http.Request) *jwt.Token {
	tokenString := extractToken(r)
	if tokenString == "" {
		panic(service_error.UnauthorizedError{})
	}

	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		log.Print(err)
		panic(service_error.UnauthorizedError{})
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		panic(service_error.UnauthorizedError{})
	}
	return token
}

func GetUser(r *http.Request) string {
	token := VerifyToken(r)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		panic(errors.New("could not parse token to get user"))
	}

	user, ok := claims["user"].(string)
	if !ok {
		panic(errors.New("could not parse token to get user"))
	}

	return user
}
