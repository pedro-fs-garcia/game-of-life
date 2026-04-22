BIN_DIR := bin
EXERCISE ?= gameoflife
BINARY := $(BIN_DIR)/$(EXERCISE)

# Map exercise to its package path
PKG_gameoflife         := ./cmd/app
PKG_langtons-ant       := ./cmd/langtonsant
PKG_cellular-automata  := ./cmd/cellularautomata
PKG_wireworld          := ./cmd/wireworld
PKG_brians-brain       := ./cmd/brians-brain
PKG_seeds-automaton    := .
PKG := $(PKG_$(EXERCISE))

.PHONY: build run tidy fmt test vet clean check

build:
	cd $(EXERCISE) && \
	go build -o ../$(BINARY) $(PKG)

run:
	cd $(EXERCISE) && \
	go run $(PKG)

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
