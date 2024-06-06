.PHONY: compose create-db drop-db migrate-up migrate-down start

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

start:
	go run app/api/main.go
