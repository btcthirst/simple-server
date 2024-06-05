FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM scratch

COPY --from=builder /app/main /app/main

EXPOSE 8081

CMD ["/app/main"]