package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	pb "house1004/services/captcha/proto"
	"log"
	"net/http"
	"time"
)

func GetCaptcha(ctx *gin.Context) {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := pb.NewGetCaptchaClient(conn)
	resp, err := cli.Call(ctx, &pb.CapRequest{})
	if resp == nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx.Writer.Header().Set("Content-Type", "image/png")
	http.ServeContent(ctx.Writer, &http.Request{}, "a.png", time.Time{}, bytes.NewReader(resp.ImgData))
}
