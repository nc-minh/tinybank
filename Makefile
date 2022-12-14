postgres:
	docker run --name postgres14 -p 5555:5432 -e POSTGRES_USER=mars -e POSTGRES_PASSWORD=mars -d postgres:14

createdb:
	docker exec -it postgres14 createdb --username=mars --owner=mars tinybank

dropdb:
	docker exec -it postgres14 dropdb --username=mars tinybank

migrateup:
	migrate -path db/migrations -database "postgresql://mars:mars@localhost:5555/tinybank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://mars:mars@localhost:5555/tinybank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test