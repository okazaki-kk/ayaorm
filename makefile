fmt:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...

test-pretty:
	set -o pipefail && go test -v ./... fmt -json | tparse -all

lint:
	golangci-lint run ./...

build:
	go build -o ayaorm ./cmd/ayaorm/

testgen:
	make build && cd tests/sqlite/ && cp ../../ayaorm . && ./ayaorm schema.go
	cd tests/mysql/ && cp ../../ayaorm . && ./ayaorm schema.go

test-pretty:
	set -o pipefail && go test -v ./... fmt -json | tparse -all
