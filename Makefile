SRC_PATHS := ./cmd ./pkg
SRC := $(shell find $(SRC_PATHS) -name '*.go')

all: scrumkin

scrumkin: $(SRC)
	go install -v cmd/$@/$@.go

test: $(SRC)
	go test $(addsuffix '/...', $(SRC_PATHS))

vendor: $(SRC)
	go get -v github.com/tools/godep
	go get -d -v ./...
	godep save ./...

.PHONY: all scrumkin test vendor
