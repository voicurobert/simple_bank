postgres:
	docker run --name postgres12 -p 5437:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

create_db:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgres12 dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5437/simple_bank?sslmode=disable" -verbose up

migrate_up1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5437/simple_bank?sslmode=disable" -verbose up 1

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5437/simple_bank?sslmode=disable" -verbose down

migrate_down1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5437/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/voicurobert/simple_bank/db/sqlc Store

.PHONY: postgres create_db drop_db migrate_up migrate_up1 migrate_down migrate_down1 server