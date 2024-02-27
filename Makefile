all: client server

.PHONY: client
client:
	go build ./cmd/client

.PHONY: server
server:
	go build ./cmd/server

.PHONY: clean
clean:
	rm -f client server

.PHONY: client
mocks:
	go generate ./...
