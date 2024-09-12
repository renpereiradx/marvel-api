package helpers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/renpereiradx/marvel-api/model"
	"github.com/renpereiradx/marvel-api/server"
)

func GetJWTSecret(server server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &model.AppClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(server.Config().JwtSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	return token, err
}
