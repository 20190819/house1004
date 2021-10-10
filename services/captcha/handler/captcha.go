package handler

import (
	"bytes"
	"context"
	pb "house1004/services/captcha/proto"

	"github.com/dchest/captcha"
)

const stdWidth=90
const stdHeight=40

type Captcha struct{}

func (e *Captcha) Code(ctx context.Context, req *pb.CapRequest) (*pb.CapResponse, error) {

	var content bytes.Buffer
	id := captcha.NewLen(4)
	captcha.WriteImage(&content, id, stdWidth, stdHeight)
	return &pb.CapResponse{ImgId: id}, nil
}

func (e *Captcha) Img(ctx context.Context, req *pb.ImgRequest) (*pb.ImgResponse, error) {
	var content bytes.Buffer
	captcha.WriteImage(&content, req.Id, captcha.StdWidth, captcha.StdHeight)
	return &pb.ImgResponse{Data: content.Bytes()}, nil
}

func (e *Captcha) Validate(ctx context.Context, req *pb.VRequest) (*pb.VResponse, error) {

	res := captcha.VerifyString(req.Id, req.Value)
	// fmt.Println("server VerifyString res>>>", res)
	return &pb.VResponse{
		Res: res,
	}, nil
}
