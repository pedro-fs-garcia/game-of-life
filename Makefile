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
	./bin/gameoflife

clean:
	rm -rf $(BIN)

check: fmt vet test