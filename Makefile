all: client server

.PHONY: client
client: generate build_client 

.PHONY: build_client
build_client:
	go build ./cmd/client

.PHONY: server
server: generate build_server 

.PHONY: build_server
build_server:
	go build ./cmd/server

.PHONY: clean
clean:
	rm -f client server

.PHONY: test
test: generate run_tests

.PHONY: run_tests
run_tests:
	go test ./...

.PHONY: generate
generate:
	go generate ./...
