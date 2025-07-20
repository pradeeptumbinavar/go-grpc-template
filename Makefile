GO_PKG = go-grpc-template
PROTO_SRC = proto/user.proto

.PHONY: build run docker proto

proto:
	@protoc --go_out=. --go-grpc_out=. $(PROTO_SRC)

build:
	@go build -o bin/server ./cmd/server

run: build
	@PORT=50051 ./bin/server

docker:
	@docker build -t go-grpc-template:latest .
