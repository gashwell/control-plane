.PHONY: run build test clean

run:
	go run cmd/server/main.go

build:
	go build -o control-plane cmd/server/main.go

test:
	go test ./...

clean:
	rm -f control-plane

install:
	go mod download
