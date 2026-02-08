fmt:
	gofmt -w .

test:
	go test ./...

build:
BIN := bin
BINARY := $(BIN)/gameoflife

.PHONY: all fmt test build run vet clean mod-tidy check

all: build

build: 
	mkdir -p $(BIN)
	go build -o $(BINARY) ./cmd/app

vet:
	go vet ./...

run:
	go run ./cmd/app

clean:
	rm -rf $(BIN)

check: fmt vet test