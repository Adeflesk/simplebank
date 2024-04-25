postgres:
	docker run --name postgres16.2 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-bullseye

createdb:
	docker exec -it postgres16.2 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16.2 dropdb simple_bank

migrateup:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose up
	
migrateup1:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover  ./...

server:
	go run main.go
	
mock: 
	mockgen -package mockDb -destination db/mock/store.go github.com/Adeflesk/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock

