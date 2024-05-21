This is based on a youtube series https://youtu.be/rx6CPDK_5mU?si=_gupEzl0PzUwhWrx

The Diagram for the postgress database was made on https://dbdiagram.io/

docker pull postgres:16.2-alpine

docker run --name postgres16.2 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-bullseye

docker exec -it postgres16.2 psql -U root

docker inspect \
  -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres16.2

  migrate create -ext sql -dir db/migration -seq init_schema // mkdir -p db/migration

  Make file  add the commands, make sure of *Tabs instead of spaces 

Running the migrate command to create the initial schema
  migrate -path db/migration --database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose up

Install sqlc 'brew install sqlc"

ACID
Atomicity
  Either all operations complete successfully or none do.
Consistency
  The db state must be valid after the transaction. All constraints must be satisfied
Isolation
   Concurrent transactions cannot affect each other

Durability
   Data written by a sucessful transaction must be permanent  

Completed all the tests of CRUD and ACID

Set up github actions with creating .github/workflows/ci.yaml file

Setup api using golang and gin  
// https://github.com/gin-gonic/gin


Used viper for reading app.env and loading a config file
// https://github.com/spf13/viper

Mocking database with mockgen 
https://github.com/uber-go/mock

To generate the mockdb use command make mock


To create branch to move on with git checkout -b ft/docker

Create the Docker file, making sure that you create a small image by building in stages

to remove built images  use Docker rmi "image id"
to stop built running containers use Docker stop "container id"
to remove built containers use Docker rm "container id"

to start running containers use
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@172.17.0.2/simple_bank?sslmode=disable" simplebank:latest