package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello %s", task)
	} else if r.Method == http.MethodPost {
		dec := json.NewDecoder(r.Body)

		var m requestBody
		if err := dec.Decode(&m); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		task = m.Message
		fmt.Fprint(w, "success")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", taskHandler).Methods("GET", "POST")

	http.ListenAndServe("localhost:8080", router)
}
