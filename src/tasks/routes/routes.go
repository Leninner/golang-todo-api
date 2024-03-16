package tasks

import (
	"encoding/json"
	"net/http"
	"todo-api/src/core/database"
	"todo-api/src/core/utils"
	"todo-api/src/tasks"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", DeleteUserHandler).Methods("DELETE")
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []tasks.Task{}
	database.DB.Find(&tasks)
	response := utils.NewResponseMessage("Tasks retrieved", tasks)
	json.NewEncoder(w).Encode(response)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.NewResponseMessage("Get task", nil))
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := database.DB.Create(&task)

	if createdTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Error creating task", nil))
	}

	json.NewEncoder(w).Encode(utils.NewResponseMessage("Task created", &task))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
