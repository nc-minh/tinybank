DB_URL=postgresql://mars:mars@localhost:5555/tinybank?sslmode=disable
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
	migrate -path db/migrations -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path db/migrations -database "${DB_URL}" -verbose up 1

migratedown:
	migrate -path db/migrations -database "${DB_URL}" -verbose down

migratedown1:
	migrate -path db/migrations -database "${DB_URL}" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nc-minh/tinybank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/nc-minh/tinybank/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=tinybank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc -f

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis --network bank-network -p 6379:6379 -d redis:7-alpine

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration sqlc test server mock proto evans redis db_docs db_schema
