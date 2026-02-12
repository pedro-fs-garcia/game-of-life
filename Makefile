BIN_DIR := bin
EXERCISE ?= gameoflife
BINARY := $(BIN_DIR)/$(EXERCISE)

.PHONY: build run tidy fmt test vet clean check

build:
	cd $(EXERCISE) && \
	go build -o ../$(BINARY) ./cmd/app

run:
	cd $(EXERCISE) && \
	go run ./cmd/app

tidy:
	cd $(EXERCISE) && \
	go mod tidy

fmt:
	cd $(EXERCISE) && \
	gofmt -w .

test:
	cd $(EXERCISE) && \
	go test ./...

vet:
	cd $(EXERCISE) && \
	go vet ./...

clean:
	rm -rf $(BIN_DIR)

check: fmt vet test
