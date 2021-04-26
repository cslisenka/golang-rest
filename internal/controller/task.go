package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"slisenko.com/kslisenko/golang-rest/internal/model"
	"slisenko.com/kslisenko/golang-rest/internal/service"
)

type Controller struct {
	Repo *service.TaskService
}

// Controller/transport layer
func (c *Controller) GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("getTasks called")

	tasks, _ := (*c.Repo).GetTasks()

	w.Header().Set("Content-Type", "application/json")
	error := json.NewEncoder(w).Encode(tasks)
	if error != nil {
		log.Println(error)
	}
}

func (c *Controller) GetTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	log.Println("getTaskById called", id)
	w.Header().Set("Content-Type", "application/json")

	task, err := (*c.Repo).GetTaskById(id)

	if err != nil {
		log.Println("error getting task", err)
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(task)
}

func (c *Controller) AddTask(w http.ResponseWriter, r *http.Request) {
	log.Println("addTask called")
	var task model.Task

	error := json.NewDecoder(r.Body).Decode(&task)
	if error != nil {
		log.Println("Error parsing request JSON", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	(*c.Repo).AddTask(&task)
	log.Println("Added task", task)
}
