BIN_DIR := bin
EXERCISE ?= gameoflife
BINARY := $(BIN_DIR)/$(EXERCISE)

# Map exercise to its cmd directory
CMD_gameoflife := app
CMD_langtons-ant := langtonsant
CMD_cellular-automata := cellularautomata
cmd_wireworld := wireworld
CMD_DIR := $(CMD_$(EXERCISE))

.PHONY: build run tidy fmt test vet clean check

build:
	cd $(EXERCISE) && \
	go build -o ../$(BINARY) ./cmd/$(CMD_DIR)

run:
	cd $(EXERCISE) && \
	go run ./cmd/$(CMD_DIR)

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
