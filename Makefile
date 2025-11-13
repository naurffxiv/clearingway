.PHONY: test build format lint lint-fix

test:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out

clearingway:
	go build -o clearingway

postgres:
	docker-compose up postgres_local

format:
	goimports -w .

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix