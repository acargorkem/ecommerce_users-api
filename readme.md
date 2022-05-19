# Microservice of restful users api
In the application, CRUD operations implemented with restful MVC pattern. Used Go and PostgreSQL.   

# Installation

## Clone the repository

> git clone https://github.com/acargorkem/ecommerce_users-api

Create a file named <strong>“.env”</strong> and fill it like “.env example” file for environment variables.

## Run with Docker
### Prerequisites:
[Docker](https://docs.docker.com/get-docker/)

Run the application with docker compose

> docker-compose up

## Run without Docker

### Prerequisites
[Go](https://go.dev/doc/install)

[PostgreSQL](https://www.postgresql.org/download/)

[golang-migrate/migrate](https://github.com/golang-migrate/migrate)

Migrate your database

>make migrateup

Start web application

>go run main.go

After that you can access the api at <strong>localhost:8080</strong>
