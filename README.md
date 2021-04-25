# golang-rest
Sample REST API application on golang

## cmd
Contains executable files

## internal
Contains packages for internal usage. It is not possible to import it to other projects.

## how to build and run
docker build -t golang-rest:1.0 .
docker run -p 9090:9090 --name golang-rest golang-rest:1.0