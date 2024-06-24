build:
	@go build -o bin/chat-app cmd/main.go
test:
	@go test -v ./...
run: build
	@./bin/chat-app