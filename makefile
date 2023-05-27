fmt:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...

lint:
	golangci-lint run ./...

build:
	go build -o ayaorm ./cmd/ayaorm/
