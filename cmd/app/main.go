package main

import (
	"log"
	"net/http"
	"todo-api/src/auth"
	"todo-api/src/core/database"
	"todo-api/src/core/utils"
	tasks "todo-api/src/tasks/routes"
	users "todo-api/src/users/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.GetConnection()
	err := database.SetupModels()

	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	StartServer()
}

func StartServer() {
	var router = mux.NewRouter()

	SetupMiddlewares(router)
	tasks.SetupRoutes(router)
	users.SetupRoutes(router)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router), "Server failed")
}

func SetupMiddlewares(router *mux.Router) {
	authService := auth.NewJWTAuthService(&auth.Configuration{
		JwtSecret:     "JWT-supersecret-sign-password",
		JwtExpiration: "15m",
	})

	router.Use(utils.LogRequestMiddleware, utils.HandleExceptionMiddleware, utils.ApplicationJSONMiddleware, authService.AuthMiddleware)
}
