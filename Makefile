DOCKER_REPO := kobtea/sample-circleci

PHONY: setup lint fmt test build release docker-build docker-release

setup:
	go get golang.org/x/tools/cmd/goimports

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.27.0 golangci-lint run -E golint,gofmt,goimports

fmt:
	go fmt ./...
	goimports -l -w .

test:
	go test ./...

build:
	go build -ldflags='-s -w -X github.com/kobtea/sample-circleci/cmd.Version=$(shell cat VERSION)'
