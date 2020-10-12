.PHONY: build
build:
	@go build -o go-todo .

.PHONY: dev
dev:
	@go run ./server.go

.PHONY: test
test:
	@go test ./...
