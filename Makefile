postgres:
	docker run --name postgres14 --network bank-network -p 5555:5432 -e POSTGRES_USER=mars -e POSTGRES_PASSWORD=mars -d postgres:14

start-postgres:
	docker start postgres14

stop-postgres:
	docker stop postgres14

createdb:
	docker exec -it postgres14 createdb --username=mars --owner=mars tinybank

dropdb:
	docker exec -it postgres14 dropdb --username=mars tinybank

migrateup:
	migrate -path db/migrations -database "postgresql://mars:mars@localhost:5555/tinybank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://mars:mars@localhost:5555/tinybank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://mars:mars@localhost:5555/tinybank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://mars:mars@localhost:5555/tinybank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nc-minh/tinybank/db/sqlc Store

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock proto evans