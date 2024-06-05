PHONY: build
build:
	go build -o ./bin/api ./cmd/api/main.go
PHONY: run
run: build
	./bin/api
PHONY: dbuild
dbuild:
	docker build -t my-golang-app .
PHONY: drun
drun:
	docker run -p 8081:8081 my-golang-app