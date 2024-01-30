build:
	docker-compose build golang-test-restapi

run:
	docker-compose up golang-test-restapi

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go