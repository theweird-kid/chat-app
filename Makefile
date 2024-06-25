build:
	@go build -o bin/chat-app cmd/main.go
test:
	@go test -v ./...
run: build
	@./bin/chat-app
migrate:
	@cd ./cmd/migrations/schema && goose postgres postgres://postgres:postgres@localhost:5432/chat_db up