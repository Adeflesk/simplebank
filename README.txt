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

To add new tables to the schema 
run command migrate create -ext sql -dir db/migration -seq add_users

Update tests for new tables
Update the mockDb with make mock