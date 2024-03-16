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
	database.GetDatabaseConnection()
	err := database.MigrateModels()

	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	StartServer()
}

func StartServer() {
	var router = mux.NewRouter()

	router.Use(utils.LogRequestMiddleware, utils.HandleExceptionMiddleware)

	router.HandleFunc("/tasks", tasks.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", tasks.GetTaskHandler).Methods("GET")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router), "Server failed")
}
