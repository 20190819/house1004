package main

import (
	"fmt"
	"house1004/bootstrap"
	"house1004/exceptions"
	"house1004/services"
	"house1004/services/user/handler"
	"house1004/services/user/models/user"
	pb "house1004/services/user/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	bootstrap.LoadEnv("./../../")
	bootstrap.ConnectDB()
	user.Migration()

	lis, err := net.Listen("tcp", services.SrvAddressUser)
	exceptions.Fatal(err)

	gsrv := grpc.NewServer()
	pb.RegisterUserServer(gsrv, handler.User{})

	go gsrv.Serve(lis)
	fmt.Printf("start user server%s\n", services.SrvAddressUser)
	select {}
}
