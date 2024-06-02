PHONY: build
build:
	go build -o ./bin/api ./cmd/api/main.go
PHONY: run
run: build
	./bin/api