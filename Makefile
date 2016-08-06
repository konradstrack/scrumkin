SRC_PATHS := ./cmd ./pkg
SRC := $(shell find $(SRC_PATHS) -name '*.go')

all: scrumway

scrumway: $(SRC)
	go install -v cmd/$@/$@.go

test: $(SRC)
	go test $(SRC_PATHS)

vendor: $(SRC)
	go get -v github.com/tools/godep
	go get -d -v ./...
	godep save ./...

.PHONY: all scrumway test vendor