package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"slisenko.com/kslisenko/golang-rest/internal/controller"
	"slisenko.com/kslisenko/golang-rest/internal/service"
)

func main() {
	// cfg := elasticsearch.Config{
	// 	Addresses: []string{
	// 		"http://localhost:9200",
	// 	},
	// }

	// es, err := elasticsearch.NewClient(cfg)
	// log.Println("Created elasticsearch client", es)

	// if err != nil {
	// 	log.Fatalf("Error creating elasticsearch client %s", err)
	// 	return
	// }

	// info, err := es.Nodes.Info()
	// log.Println("Calling info endpoint")
	// // TODO why should we close it explicitly?
	// // defer info.Body.Close()
	// if err != nil {
	// 	log.Fatalf("Error calling elastic %s", err)
	// }
	// log.Println(info.Status())
	// defer info.Body.Close()

	// // Index sample document
	// newTask := model.Task{
	// 	ID:          "4",
	// 	Description: "Sample task 4",
	// }

	// newTaskJson, _ := json.Marshal(newTask)

	// indexReq := esapi.IndexRequest{
	// 	Index:      "tasks",
	// 	DocumentID: newTask.ID,
	// 	Body:       bytes.NewReader(newTaskJson),
	// 	Refresh:    "true",
	// }

	// indexResp, err := indexReq.Do(context.Background(), es)
	// if err != nil || indexResp.IsError() {
	// 	log.Fatalf("Error indexing document %s", err)
	// 	return
	// }

	// log.Println(indexResp)

	// defer indexResp.Body.Close()

	// // Find document by ID
	// findByIdReq := esapi.GetRequest{
	// 	Index:      "tasks",
	// 	DocumentID: newTask.ID,
	// }
	// findByIdResp, err := findByIdReq.Do(context.Background(), es)
	// if err != nil || findByIdResp.IsError() {
	// 	log.Fatalf("Error finding document %s %s", err, findByIdResp)
	// }

	// log.Println(findByIdResp)
	// defer findByIdResp.Body.Close()

	// // Find all documents
	// searchReq := esapi.SearchRequest{
	// 	Index: []string{"tasks"},
	// }

	// searchResp, err := searchReq.Do(context.Background(), es)
	// if err != nil {
	// 	log.Fatalf("Error searching documents %s", err)
	// }
	// log.Println(searchResp)
	// // Decode json, create structure for JSON parsing
	// // Create structure that match JSON output
	// type ElasticResponse struct {
	// 	Hits struct {
	// 		Hits []struct {
	// 			Source struct {
	// 				Id          string `json:"id"`
	// 				Description string `json:"description"`
	// 			} `json:"_source"`
	// 		} `json:"hits"`
	// 	} `json:"hits"`
	// }

	// var parsedResponse ElasticResponse
	// json.NewDecoder(searchResp.Body).Decode(&parsedResponse)

	// log.Println(parsedResponse.Hits.Hits)

	// defer searchResp.Body.Close()

	var taskService service.TaskService = &service.InMemoryTaskRepository{}
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
