package users

import (
	"encoding/json"
	"net/http"
	"todo-api/src/core/database"
	"todo-api/src/core/utils"
	"todo-api/src/tasks"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []tasks.Task{}
	database.DB.Find(&tasks)
	response := utils.NewResponseMessage("Tasks retrieved", tasks)
	json.NewEncoder(w).Encode(response)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := database.DB.Create(&task)

	if createdTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating task"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
