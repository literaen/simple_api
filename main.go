package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var task []Task
	tasks := DB.Find(&task)

	if tasks.Error != nil {
		http.Error(w, tasks.Error.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var task Task
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedTask := Task{
		Task:   task.Task,
		IsDone: task.IsDone,
	}

	result := DB.Model(&Task{}).Where("id = ?", taskID).Updates(updatedTask)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var newTask Task
	if err := DB.First(&newTask, taskID).Error; err != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Обновляем задачу с ID: %d. Новые данные: Task=%s, IsDone=%v\n", taskID, updatedTask.Task, updatedTask.IsDone)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var task Task

	result := DB.Where("id = ?", taskID).Delete(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", UpdateTask).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteTask).Methods("DELETE")

	http.ListenAndServe("localhost:8080", router)
}
