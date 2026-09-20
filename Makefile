build:
	@go build -o bin/goApiJwt cmd/main.go

test:
	@go test ./... -v

run:
	@go run cmd/main.go

migration:
	@go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir cmd/migrate/migrations -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down