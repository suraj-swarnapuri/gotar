.PHONY: build
build:
	go build -o bin/gotar-cli cmd/gotar-cli/main.go

.PHONY: gotar
gotar:
	go run cmd/gotar/main.go