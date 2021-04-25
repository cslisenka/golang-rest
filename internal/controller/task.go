package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"slisenko.com/kslisenko/golang-rest/internal/model"
	"slisenko.com/kslisenko/golang-rest/internal/service"
)

// Controller/transport layer
func GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("getTasks called")

	tasks := service.GetTasks()

	w.Header().Set("Content-Type", "application/json")
	error := json.NewEncoder(w).Encode(tasks)
	if error != nil {
		log.Fatal(error)
	}
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	log.Println("getTaskById called", id)
	w.Header().Set("Content-Type", "application/json")

	task, err := service.GetTaskById(id)

	if err != nil {
		log.Fatal("error getting task", err)
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(task)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	log.Println("addTask called")
	var task model.Task

	error := json.NewDecoder(r.Body).Decode(&task)
	if error != nil {
		log.Fatal("Error parsing request JSON", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	service.AddTask(&task)
	log.Println("Added task", task)
}
