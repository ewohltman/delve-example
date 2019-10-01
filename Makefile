.PHONY: lint test debug build

lint:
	golangci-lint run --enable-all --deadline=5m ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

debug:
	dlv debug delve-example.go

build:
	go build delve-example.go
