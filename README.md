# Book Store

Book Store is a service for online book store.

## Onboarding and Development Guide

### Dependencies

- Git
- Docker
- Go 1.22
- Golang Migrate (For database migration. See [golang-migrate installation](https://github.com/golang-migrate/migrate))
- psql (For database seed. See [PostgreSQL installation](https://www.postgresql.org/download))
- Swag (For generating API docs with Swagger. See [swag installation](https://github.com/swaggo/swag))

### Installation & Getting started

1. Clone this repository and install the prerequisites above
2. Copy `.env` from `env.sample` and modify the configuration value appropriately
3. Run the container `make compose`
4. Setup the database `make create-db && make migrate-up && make seed`
5. Download modules `go mod download`
6. Run the service by running `make start`
7. Go to API Docs http://localhost:9999/swagger/index.html
