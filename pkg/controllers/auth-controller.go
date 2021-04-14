package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthController struct {
	ExpiresInMinutes time.Duration
}

func (controller *AuthController) GetToken(w http.ResponseWriter, r *http.Request) {
	username, _, _ := r.BasicAuth()
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * controller.ExpiresInMinutes).Unix(),
		Issuer:    "k8ssandra-api-service",
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString([]byte("secret"))
	w.Write([]byte(jwtToken))
}
