# Build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz   

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder  /app/migrate.linux-amd64 ./migrate
COPY start.sh .
COPY wait-for.sh .
COPY datasources/postgresql/migration ./migration

EXPOSE 8081
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
