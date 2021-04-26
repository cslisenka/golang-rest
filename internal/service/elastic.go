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
	Host      string
	IndexName string
	es        *elasticsearch.Client
}

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

type elasticGetResponse struct {
	Source struct {
		Id          string `json:"id"`
		Description string `json:"description"`
	} `json:"_source"`
}

func (tasks *ElasticSearchTaskRepository) Init() error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			tasks.Host,
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
		log.Println("Error calling elastic", err)
		return err
	}
	log.Printf("Getting elasticsearch node info, status %s", info.Status())

	return nil
}

func (tasks *ElasticSearchTaskRepository) InitSampleData() error {
	err1 := tasks.AddTask(&model.Task{ID: "1", Description: "Task 1"})
	err2 := tasks.AddTask(&model.Task{ID: "2", Description: "Task 2"})

	if err1 != nil || err2 != nil {
		return errors.New(err1.Error() + err2.Error())
	}

	return nil
}

func (tasks *ElasticSearchTaskRepository) GetTasks() ([]model.Task, error) {
	// Find all documents
	searchReq := esapi.SearchRequest{
		Index: []string{tasks.IndexName},
	}

	searchResp, err := searchReq.Do(context.Background(), tasks.es)

	if err != nil {
		log.Println("Error searching documents", err)
	}

	defer searchResp.Body.Close()

	log.Println("elastic search response", searchResp)

	return parseElasticSearchResponse(searchResp)
}

func (tasks *ElasticSearchTaskRepository) AddTask(task *model.Task) error {
	newTaskJson, _ := json.Marshal(task)

	indexReq := esapi.IndexRequest{
		Index:      tasks.IndexName,
		DocumentID: task.ID,
		Body:       bytes.NewReader(newTaskJson),
		Refresh:    "true",
	}

	indexResp, err := indexReq.Do(context.Background(), tasks.es)

	if err != nil || indexResp.IsError() {
		log.Println("Error indexing document", err)
		return err
	}

	defer indexResp.Body.Close()

	log.Println("elastic add task response", indexResp)

	return nil
}

func (tasks *ElasticSearchTaskRepository) GetTaskById(id string) (model.Task, error) {
	findByIdReq := esapi.GetRequest{
		Index:      tasks.IndexName,
		DocumentID: id,
	}

	findByIdResp, err := findByIdReq.Do(context.Background(), tasks.es)
	if err != nil || findByIdResp.IsError() {
		log.Println("Error finding document", err, findByIdResp)
	}

	defer findByIdResp.Body.Close()

	log.Println("elastic get by ID response", findByIdResp)

	return parseElasticGetResponse(findByIdResp)
}

func parseElasticSearchResponse(resp *esapi.Response) ([]model.Task, error) {
	var parsedResponse elasticResponse
	err := json.NewDecoder(resp.Body).Decode(&parsedResponse)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task = []model.Task{}
	for _, hit := range parsedResponse.Hits.Hits {
		tasks = append(tasks, model.Task{ID: hit.Source.Id, Description: hit.Source.Description})
	}
	return tasks, nil
}

func parseElasticGetResponse(resp *esapi.Response) (model.Task, error) {
	var parsedResponse elasticGetResponse
	err := json.NewDecoder(resp.Body).Decode(&parsedResponse)
	if err != nil {
		return model.Task{}, err
	}

	return model.Task{
		ID:          parsedResponse.Source.Id,
		Description: parsedResponse.Source.Description,
	}, nil
}
