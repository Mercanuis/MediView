GOPATH := $(shell go env GOPATH)

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: coverage
coverage:
	@go test -race -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: main
main:
	@go run ./cmd/main/main.go