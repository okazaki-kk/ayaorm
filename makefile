fmt:
	go fmt ./...
test:
	go test -v ./... ./_test

lint:
	golangci-lint run ./...

build:
	go build -o ayaorm ./cmd/ayaorm/
