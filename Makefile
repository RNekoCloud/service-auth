postgres:
	docker run --name=cvz_auth  -p 5434:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine

proto:
	protoc --go_out=. --go_opt=paths=source_relative \ --go-grpc_out=. --go-grpc_opt=paths=source_relative \ internal/proto/auth.proto

createdb:
	docker exec -it cvz_auth createdb --username=root --owner=root cvz_auth

dropdb:
	docker exec -it cvz_auth dropdb cvz_auth

migrateup:
	migrate -path internal/db/migration -database "postgres://root:root@127.0.0.1:5434/cvz_auth?sslmode=disable" -verbose up

.PHONY: postgres proto createdb dropdb migrateup
