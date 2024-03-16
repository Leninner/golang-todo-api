package main

import (
	"log"
	"net/http"
	"todo-api/src/core/database"
	"todo-api/src/core/utils"
	tasks "todo-api/src/tasks/routes"

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

	log.Fatal(http.ListenAndServe(":8080", router), "Server failed")
	log.Println("Server started on port 8080")
}

func SetupMiddlewares(router *mux.Router) {
	router.Use(utils.LogRequestMiddleware, utils.HandleExceptionMiddleware)
}
