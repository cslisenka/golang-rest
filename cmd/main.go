package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"slisenko.com/kslisenko/golang-rest/internal/controller"
	"slisenko.com/kslisenko/golang-rest/internal/service"
)

func main() {
	// Getting elastic host/port from env variables
	elasticHost, elasticPort := os.Getenv("ELASTIC_HOST"), os.Getenv("ELASTIC_PORT")
	elasticURL := fmt.Sprintf("http://%s:%s", elasticHost, elasticPort)
	log.Println("Got elastic host/port", elasticURL)

	//var taskService service.TaskService = &service.InMemoryTaskRepository{}
	var taskService service.TaskService = &service.ElasticSearchTaskRepository{
		Host:      elasticURL,
		IndexName: "tasks",
	}
	taskService.Init()
	taskService.InitSampleData()

	controller := controller.Controller{Repo: &taskService}

	r := mux.NewRouter()
	// Define routes
	r.HandleFunc("/task", controller.GetTasks).Methods("GET")
	r.HandleFunc("/task", controller.AddTask).Methods("POST")
	r.HandleFunc("/task/{id}", controller.GetTaskById).Methods("GET")
	// TODO add POST/DELETE

	log.Println("Starting server on port 9090")
	// TODO extract port as command line argument
	error := http.ListenAndServe(":9090", r)
	if error != nil {
		log.Fatal(error)
	}
}
