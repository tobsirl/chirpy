BIN=./bin/out

.PHONY: run build

build:
	go build -o $(BIN)

run: build
	$(BIN)
