package users

import (
	"encoding/json"
	"net/http"
	"time"
	"todo-api/src/core/database"
	"todo-api/src/core/utils"
	"todo-api/src/users"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginDTO LoginDTO

	if err := json.NewDecoder(r.Body).Decode(&loginDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid request", nil))
		return
	}

	token, err := Login(loginDTO.Email, loginDTO.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid email or password", nil))
		return
	}

	json.NewEncoder(w).Encode(utils.NewResponseMessage("Login successful", token))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user users.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid request", nil))
		return
	}

	errors := utils.ValidateStruct(user)

	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid request", errors))
		return
	}

	if hasBeenCreated := database.DB.Create(&user); hasBeenCreated.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Error creating user", nil))
		return
	}

	json.NewEncoder(w).Encode(utils.NewResponseMessage("User created", user))
}

type CustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.MapClaims
}

func Login(email, password string) (*LoginResponseDTO, error) {
	user := &users.User{}
	database.DB.Limit(1).Find(user, "email = ? AND password = ?", email, password)

	claims := CustomClaims{
		user.ID,
		user.Email,
		user.Role,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 72).Unix(),
			"iat": time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("JWT-supersecret-sign-password"))

	if err != nil {
		return nil, err
	}

	return &LoginResponseDTO{
		Token: tokenStr,
	}, nil
}
