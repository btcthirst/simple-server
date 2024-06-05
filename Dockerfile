FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod download
 
RUN go build -o api ./cmd/api/main.go

FROM ubuntu:latest

COPY --from=builder /app/api /app/api

CMD [ "/app/api" ]