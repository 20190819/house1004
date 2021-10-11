package main

import (
	"fmt"
	"house1004/exceptions"
	"house1004/services"
	handler "house1004/services/captcha/handler"
	pb "house1004/services/captcha/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", services.SrvAddressCaptcha)
	exceptions.Fatal(err)
	gsrv := grpc.NewServer()
	pb.RegisterGetCaptchaServer(gsrv, new(handler.Captcha))

	go gsrv.Serve(lis)
	fmt.Printf("start grpc server%s\n", services.SrvAddressCaptcha)
	select {

	}
}
