package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"slisenko.com/kslisenko/golang-rest/internal/model"
)

// Service layer
type ElasticSearchTaskRepository struct {
	es *elasticsearch.Client
}

const IndexName string = "tasks"

// Internal data structures
type elasticResponse struct {
	Hits struct {
		Hits []struct {
			Source struct {
				Id          string `json:"id"`
				Description string `json:"description"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (tasks *ElasticSearchTaskRepository) Init(host string) error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			host,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating elasticsearch client %s", err)
		return err
	}

	tasks.es = es

	info, err := es.Nodes.Info()
	if err != nil {
		log.Fatalf("Error calling elastic %s", err)
	}
	log.Printf("Getting elasticsearch node info, status %s", info.Status())

	return nil
}

func (tasks *ElasticSearchTaskRepository) InitSampleData() {
	tasks.AddTask(&model.Task{ID: "1", Description: "Task 1"})
	tasks.AddTask(&model.Task{ID: "2", Description: "Task 2"})
}

func (tasks *ElasticSearchTaskRepository) GetTasks() { //[]model.Task
	// Find all documents
	searchReq := esapi.SearchRequest{
		Index: []string{IndexName},
	}

	searchResp, err := searchReq.Do(context.Background(), tasks.es)

	if err != nil {
		log.Fatalf("Error searching documents %s", err)
	}

	defer searchResp.Body.Close()

	log.Println(searchResp)
	// Decode json, create structure for JSON parsing
	// Create structure that match JSON output

	parseElasticSearchResponse(searchResp)

	// log.Println(parsedResponse.Hits.Hits)

}

func (tasks *ElasticSearchTaskRepository) AddTask(task *model.Task) {
	newTaskJson, _ := json.Marshal(task)

	indexReq := esapi.IndexRequest{
		Index:      IndexName,
		DocumentID: task.ID,
		Body:       bytes.NewReader(newTaskJson),
		Refresh:    "true",
	}

	indexResp, err := indexReq.Do(context.Background(), tasks.es)

	if err != nil || indexResp.IsError() {
		log.Fatalf("Error indexing document %s", err)
		return
	}

	defer indexResp.Body.Close()

	log.Println(indexResp)
}

func (tasks *ElasticSearchTaskRepository) GetTaskById(id string) (model.Task, error) {
	findByIdReq := esapi.GetRequest{
		Index:      IndexName,
		DocumentID: id,
	}

	findByIdResp, err := findByIdReq.Do(context.Background(), tasks.es)
	if err != nil || findByIdResp.IsError() {
		log.Fatalf("Error finding document %s %s", err, findByIdResp)
	}

	defer findByIdResp.Body.Close()

	log.Println(findByIdResp)

	// TODO parse response

	return model.Task{}, errors.New("task not found")
}

// TODO methods for response parsing
func parseElasticSearchResponse(resp *esapi.Response) { //[]model.Task
	var parsedResponse elasticResponse
	json.NewDecoder(resp.Body).Decode(&parsedResponse)

	log.Println(parsedResponse)
}
