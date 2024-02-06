build:
	docker-compose build test-crud-api

run:
	docker-compose up test-crud-api

migrate:
	migrate -path ./schema -database 'postgresql://postgres:qwerty@0.0.0.0:5433/postgres?sslmode=disable' up
