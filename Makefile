clientserver: client server

client:
	go build ./cmd/client

server:
	go build ./cmd/server

clean:
	rm -f client server

mocks:
	go generate ./...
