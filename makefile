.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: test-pretty
test-pretty:
	set -o pipefail && go test -v ./... fmt -json | tparse -all

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build:
	go build -o ayaorm ./cmd/ayaorm/

.PHONY: testgen
testgen:
	make build && cd tests/sqlite/ && cp ../../ayaorm . && ./ayaorm schema.go
	cd tests/mysql/ && cp ../../ayaorm . && ./ayaorm schema.go
