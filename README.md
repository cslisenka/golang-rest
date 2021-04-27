# golang-rest
Sample REST API application on golang

## cmd
Contains executable files

## internal
Contains packages for internal usage. It is not possible to import it to other projects.

# installing modules
go mod github.com/elastic/go-elasticsearch/v7 
go mod github.com/gorilla/mux (fix version!)

## how to build and run
docker build -t golang-rest:1.0 .
docker run -p 9090:9090 --name golang-rest golang-rest:1.0

# elasticsearch
Run in docker: 
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.12.0
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.12.0

Check health: GET http://localhost:9200/_cat/nodes?v=true&pretty
Create index: PUT http://localhost:9200/tasks
Create document: POST http://localhost:9200/tasks/_doc/1
{
	"ID" : "1",
	"Description" : "Test"
}

Get all documents: GET http://localhost:9200/tasks/_search
Get document by ID: GET http://localhost:9200/tasks/_doc/1

# kubernetes (minikube)
using own docker daemon, we need switch to it
eval $(minikube -p minikube docker-env)

Deploy app and services:
kubectl apply -f k8s/application.yaml

Get public IP/port of the NodePort service:
minikube service --url golang-rest-service