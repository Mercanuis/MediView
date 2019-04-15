GOPATH := $(shell go env GOPATH)

$(GOPATH)/bin/dep:
	@go get github.com/golang/dep/cmd/dep

.PHONY: dep
dep: $(GOPATH)/bin/dep
	@dep ensure -v

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: coverage
coverage:
	@go test -race -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: main
main:
	@go run ./cmd/main/main.go

.PHONY: main-short
main-short:
	@go run ./cmd/main/main.go short