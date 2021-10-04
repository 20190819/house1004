.PHONY: proto
proto:
	protoc -I. --proto_path=. --go_out=plugins=grpc:. ./proto/captcha.proto

.PHONY: build
build:
	go build -o captcha *.go