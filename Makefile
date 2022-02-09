.PHONY: all
all: build test

.PHONY: download
download:
		go mod tidy && go mod download all

.PHONY: build
build:
		go build -o app cmd/server/main.go

.PHONY: clean
clean: 
		go clean
		rm -rf build

.PHONY: test
test:
		go test -v -race ./cmd/... ./pkg/...
