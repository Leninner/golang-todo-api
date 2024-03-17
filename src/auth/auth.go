package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	users "todo-api/src/users/routes"

	"github.com/golang-jwt/jwt/v5"
)

type Configuration struct {
	JwtSecret     string
	JwtExpiration string
}

type JWTAuthService struct {
	Config *Configuration
}

func NewJWTAuthService(cfg *Configuration) *JWTAuthService {
	return &JWTAuthService{
		cfg,
	}
}

func (j *JWTAuthService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			json.NewEncoder(w).Encode("unauthorized - no token given")
			return
		}

		valid, err := j.ValidateJWT(strings.Split(token, " ")[1])

		if err != nil {
			json.NewEncoder(w).Encode("unauthorized - invalid token struct")
			return
		}

		if !valid {
			log.Fatal("unauthorized - invalid token")
			json.NewEncoder(w).Encode("unauthorized - token invalid")
			return
		}

		err = j.DecodeJWT(strings.Split(token, " ")[1])
		if err != nil {
			json.NewEncoder(w).Encode("unauthorized - can't decode token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (j *JWTAuthService) DecodeJWT(token string) error {
	_, err := jwt.ParseWithClaims(token, &users.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Config.JwtSecret), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (j *JWTAuthService) ValidateJWT(token string) (bool, error) {
	tokenData, err := jwt.ParseWithClaims(token, &users.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(j.Config.JwtSecret), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := tokenData.Claims.(*users.CustomClaims); ok && tokenData.Valid {
		return true, nil
	}

	return false, nil
}
