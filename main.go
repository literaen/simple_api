package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var task []Task
	tasks := DB.Find(&task)

	if tasks.Error != nil {
		http.Error(w, tasks.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	DB.Create(&task)
	fmt.Fprint(w, "Задание успешно добавлено")
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")

	http.ListenAndServe("localhost:8080", router)
}
