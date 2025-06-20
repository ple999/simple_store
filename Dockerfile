FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:latest AS final
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=builder /app/migration_file ./migration_file
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/start.sh .


EXPOSE 2000
CMD ["/app/main"]