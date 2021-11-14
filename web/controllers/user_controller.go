package controllers

import (
	"house1004/exceptions"
	"house1004/services"
	"house1004/web/helpers"
	"net/http"

	pb "house1004/services/user/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type UserController struct {
	Conn *grpc.ClientConn
	err  error
}

func (uc *UserController) HandlerRegister(ctx *gin.Context) {
	uc.Conn, uc.err = grpc.Dial(services.SrvAddressUser, grpc.WithInsecure())
	exceptions.Fatal(uc.err)

	cli := pb.NewUserClient(uc.Conn)
	req := &pb.Request{
		Phone: ctx.PostForm("phone"),
		Captcha: ctx.PostForm("captcha"),
		Password: ctx.PostForm("password"),
		PasswordConfirm: ctx.PostForm("password_confirm"),
	}
	_, err := cli.Register(ctx,req)
	exceptions.Fatal(err)
	ctx.JSON(http.StatusOK,gin.H{
		"code":helpers.SuccessCode,
		"msg":helpers.SuccessMsg,
	})
}

func (uc *UserController) HandlerLogin(ctx *gin.Context){
	name:=ctx.PostForm("name")
	password:=ctx.PostForm("password")

	uc.Conn, uc.err = grpc.Dial(services.SrvAddressUser, grpc.WithInsecure())
	exceptions.Fatal(uc.err)

	cli := pb.NewUserClient(uc.Conn)
	req:=&pb.LRequest{
		Name: name,
		Password: password,
	}
	_, err := cli.Login(ctx,req)
	exceptions.Fatal(err)
	ctx.JSON(http.StatusOK,gin.H{
		"code":helpers.SuccessCode,
		"msg":helpers.SuccessMsg,
	})
}

