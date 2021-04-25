package main

import (
	"log"
	"net/http"

	//"strconv"
	"github.com/gorilla/mux"
	"slisenko.com/kslisenko/golang-rest/internal/controller"
	"slisenko.com/kslisenko/golang-rest/internal/service"
)

func main() {
	service.InitSampleData()

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
