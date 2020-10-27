.PHONY: build
build:
	@go build -o go-todo .

.PHONY: dev
dev:
	@docker-compose up -d && air

.PHONY: test
test:
	@go test ./...
