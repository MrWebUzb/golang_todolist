.PHONY: run
run:
	go run ./cmd/

.PHONY: build
build:
	go build ./cmd/

.PHONY: test
test:
	go test ./

.DEFAULT_GOAL := run