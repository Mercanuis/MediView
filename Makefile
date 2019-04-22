GOPATH := $(shell go env GOPATH)

$(GOPATH)/bin/dep:
	@go get github.com/golang/dep/cmd/dep

.PHONY: dep
dep: $(GOPATH)/bin/dep
	@go build ./...

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: coverage
coverage:
	@go test -race -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: rabbit
rabbit:
	rabbitmq-server

.PHONY: main
main:
	rabbitmq-server &
	sleep 5 && go run ./cmd/main/main.go

.PHONY: main-short
main-short:
	rabbitmq-server &
	sleep 5 && go run ./cmd/main/main.go short

.PHONY: build
build: dep
	CGO_ENABLED=0 go build -o bin/mediview ../MediView/cmd/main/

.PHONY: docker
docker: dep && build
	docker build -t mediview .