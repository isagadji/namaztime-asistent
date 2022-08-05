NAME = marusya
ifeq ($(VERSION),)
    VERSION = $(shell git rev-parse --abbrev-ref HEAD)
endif

-include .env
export

.Pony: help
help: ## Help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.Pony: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./builds/$(NAME) ./cmd/main.go

.Pony: init
init:
	cp .env_example .env

.Pony: run
run:
	CGO_ENABLED=0 go run ./cmd/$(NAME) server

.Pony: test
test:
	CGO_ENABLED=0 go test -v ./...

.Pony: run_db
run_db:
	docker-compose up -d

.Pony: stop_db
stop_db:
	docker-compose down