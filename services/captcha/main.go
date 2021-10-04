package main

import (
	"fmt"
	handler "house1004/services/captcha/handler"
	pb "house1004/services/captcha/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const GRPC_ADDR = ":3000"

func main() {
	lis, err := net.Listen("tcp", GRPC_ADDR)
	if err != nil {
		log.Fatal(err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterGetCaptchaServer(gsrv, new(handler.Captcha))

	go gsrv.Serve(lis)
	fmt.Printf("start grpc server%s\n", GRPC_ADDR)
	for {
	}
}
