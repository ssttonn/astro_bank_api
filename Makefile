# include .env
# export $(shell sed 's/=.*//' .env)

postgres_container:
	docker run --name postgres_container -p 5432:5432 -e POSTGRES_USER=astrotify-admin -e POSTGRES_PASSWORD=sample -d postgres:latest

createdb:
	docker exec -it postgres_container createdb --username=astrotify-admin --owner=astrotify-admin astrobank_db

dropdb:
	docker exec -it postgres_container dropdb --username=astrotify-admin astrobank_db

enterdb:
	docker exec -it postgres_container psql --username=astrotify-admin astrobank_db

migrateup:
	migrate -path db/migration -database "postgresql://astrotify-admin:sample@localhost:5432/astrobank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://astrotify-admin:sample@localhost:5432/astrobank_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres_container createdb dropdb enterdb migrateup migratedown sqlc test server