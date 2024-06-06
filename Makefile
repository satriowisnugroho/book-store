.PHONY: compose create-db drop-db migrate-up migrate-down start test

compose:
	docker-compose up

create-db:
	docker exec -it book-store_db_1 createdb --username=root --owner=root book_store_development

drop-db:
	docker exec -it book-store_db_1 dropdb book_store_development

migrate-up:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/book_store_development?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:root@localhost:5433/book_store_development?sslmode=disable" -verbose down

seed:
	PGPASSWORD=root psql -U root -h localhost -p 5433 -t books -d book_store_development < db/seed.sql

start:
	go run app/api/main.go

test:
	go test -cover -coverprofile=coverage.out -json $$(go list ./... | grep -Ev "app") > ./UT-report.json

test-cover:
	make test
	go tool cover -html=coverage.out

mock-init:
	mockery --all --dir ./ --output ./test/mock --case underscore

swag-init-v1:
	swag init --pd -g internal/handler/http/v1/router.go
