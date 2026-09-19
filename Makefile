build:
	@go build -o bin/goApiJwt cmd/main.go

test:
	@go test ./... -v

run:
	@go run cmd/main.go