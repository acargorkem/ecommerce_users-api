ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run:
	go run main.go

run-docker:
	docker-compose up

build-docker:
	docker-compose up --build

stop-docker:
	docker-compose down

postgresql:
	docker-compose -f docker-compose.db.yml up

postgresql-stop:
	docker-compose -f docker-compose.db.yml down

migrateup: 
	migrate -path datasources/postgresql/migration -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -verbose up

migratedown: 
	migrate -path datasources/postgresql/migration -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -verbose down

.PHONY: run run-docker build-docker stop-docker migrateup migratedown postgresql postgresql-stop
