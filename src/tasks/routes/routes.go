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
	router.HandleFunc("/tasks/{id}", DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", UpdateTaskHandler).Methods("PUT")
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	orderBy := queries.Get("orderBy")
	sort := queries.Get("sort")

	tasks := []tasks.Task{}
	database.DB.Order(orderBy + " " + sort).Find(&tasks)

	json.NewEncoder(w).Encode(utils.NewResponseMessage("Tasks retrieved", tasks))
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := tasks.Task{}
	vars := mux.Vars(r)
	id := vars["id"]

	if !utils.IsInteger(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid task id", nil))
		return
	}

	database.DB.Limit(1).Find(&task, "id = ?", id)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Task not found", nil))
		return
	}

	json.NewEncoder(w).Encode(utils.NewResponseMessage("Task retreived", task))
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid request", nil))
		return
	}

	if errors := utils.ValidateStruct(task); errors != nil {
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Error creating task", errors))
		return
	}

	if hasCreatedTask := database.DB.Create(&task); hasCreatedTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Error creating task", nil))
	}

	json.NewEncoder(w).Encode(utils.NewResponseMessage("Task created", &task))
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if !utils.IsInteger(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid task id", nil))
		return
	}

	task := tasks.Task{}

	database.DB.Limit(1).Find(&task, "id = ?", id)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Task not found", nil))
		return
	}

	if hasBeenDeleted := database.DB.Delete(&task); hasBeenDeleted.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Error deleting task", nil))
		return
	}

	json.NewEncoder(w).Encode(utils.NewResponseMessage("Task deleted", nil))
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if !utils.IsInteger(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid task id", nil))
		return
	}

	task := tasks.Task{}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Invalid request", nil))
		return
	}

	if errors := utils.ValidateStruct(task); errors != nil {
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Error updating task", errors))
		return
	}

	database.DB.Limit(1).Select("id", "created_at", "updated_at", "deleted_at").Find(&task, "id = ?", id)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Task not found", nil))
		return
	}

	if hasBeenUpdated := database.DB.Save(&task); hasBeenUpdated.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.NewResponseMessage("Error updating task", nil))
		return
	}

	json.NewEncoder(w).Encode(utils.NewResponseMessage("Task updated", &task))
}
