tidy:
	@go mod tidy

build: tidy
	@go build -o ./bin/main ./cmd/main.go

run: build
	@./bin/main

test: tidy
	@go test ./... -v