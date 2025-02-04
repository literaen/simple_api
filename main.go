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

	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Задание успешно добавлено")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := DB.Save(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Задание %d успешно обновлено", task.ID)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := DB.Delete(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Задание %d успешно удалено", task.ID)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", UpdateTask).Methods("PATCH")
	router.HandleFunc("/api/tasks", DeleteTask).Methods("DELETE")

	http.ListenAndServe("localhost:8080", router)
}
