update:
	go get -u all

postgres:
	docker run --name go-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.15-alpine3.21

create_db:
	docker exec -it go-postgres createdb --username=root --owner=root go-postgres

drop_db:
	docker exec -it go-postgres dropdb go-postgres

migrate_up:
	 migrate -path workspace/database/migration -database "postgresql://root:secret@localhost:5432/go-postgres?sslmode=disable" -verbose up

migrate_down:
	 migrate -path workspace/database/migration -database "postgresql://root:secret@localhost:5432/go-postgres?sslmode=disable" -verbose down -all

sqlc:
	sqlc generate -f workspace/database/sqlc.yaml

test:
	go test -v -cover ./...
.PHONY: update postgres create_db drop_db migrate_up migrate_down test
