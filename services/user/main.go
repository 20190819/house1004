package main

import (
	"fmt"
	"google.golang.org/grpc"
	"house1004/bootstrap"
	"house1004/exceptions"
	"house1004/services"
	"house1004/services/user/handler"
	pb "house1004/services/user/proto"
	"net"
)

func main(){
	bootstrap.LoadEnv()
	bootstrap.ConnectDB()

	lis,err:=net.Listen("tcp",services.SrvAddressUser)
	exceptions.Fatal(err)

	gsrv:=grpc.NewServer()
	pb.RegisterUserServer(gsrv,handler.User{})

	go gsrv.Serve(lis)
	fmt.Printf("start user server%s\n", services.SrvAddressUser)
	select {

	}
}