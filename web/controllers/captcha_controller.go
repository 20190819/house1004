package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"house1004/exceptions"
	"house1004/services"
	pb "house1004/services/captcha/proto"
	"net/http"
	"time"
)

type CaptchaController struct {
	Conn *grpc.ClientConn
	err  error
}

func (cc CaptchaController) GetCaptchaCd(ctx *gin.Context) {
	cc.Conn, cc.err = grpc.Dial(services.SrvAddressCaptcha, grpc.WithInsecure())
	exceptions.Fatal(cc.err)
	cli := pb.NewGetCaptchaClient(cc.Conn)
	resp, err := cli.Code(ctx, &pb.CapRequest{})
	exceptions.Fatal(err)
	if resp == nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func (cc CaptchaController) GetCaptchaImg(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		exceptions.Fatal(exceptions.ErrParams)
	}
	cc.Conn, cc.err = grpc.Dial(":3000", grpc.WithInsecure())
	exceptions.Fatal(cc.err)
	cli := pb.NewGetCaptchaClient(cc.Conn)
	req := pb.ImgRequest{Id: id}
	resp, err := cli.Img(ctx, &req)
	exceptions.Fatal(err)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	http.ServeContent(ctx.Writer, &http.Request{}, "a.png", time.Time{}, bytes.NewReader(resp.Data))
}

func (cc CaptchaController) Verify(ctx *gin.Context) {
	cc.Conn, cc.err = grpc.Dial(":3000", grpc.WithInsecure())
	exceptions.Fatal(cc.err)
	cli := pb.NewGetCaptchaClient(cc.Conn)
	id := ctx.Param("id")
	value := ctx.PostForm("value")
	if id == "" || value == "" {
		exceptions.Fatal(exceptions.ErrParams)
	}
	fmt.Printf("id=%s,value=%s\n", id, value)
	resp, err := cli.Validate(ctx, &pb.VRequest{
		Id:    id,
		Value: value,
	})
	exceptions.Fatal(err)
	if resp.Res == true {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "验证成功",
			"data": resp.Res,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "验证失败",
			"data": resp.Res,
		})
	}

}
