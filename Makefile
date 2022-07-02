.Pony: build
build:
	GOOS=linux GOARCH=amd64 go build -o ./builds/marusya ./cmd/main.go

.Pony: init
init:
	cp .env_example .env