.PHONY: proto
proto:
	protoc -I. --proto_path=. --go_out=plugins=grpc:. ./proto/user.proto

.PHONY: build
build:
	go build -o user *.go