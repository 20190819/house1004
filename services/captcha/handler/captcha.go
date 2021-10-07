package handler

import (
	"bytes"
	"context"
	"fmt"
	pb "house1004/services/captcha/proto"

	"github.com/dchest/captcha"
)

type Captcha struct{}

func (e *Captcha) Call(ctx context.Context, req *pb.CapRequest) (*pb.CapResponse, error) {

	var content bytes.Buffer
	id := captcha.NewLen(4)
	url := fmt.Sprintf("/captcha_img/%s.png", id)
	captcha.WriteImage(&content, id, captcha.StdWidth, captcha.StdHeight)
	return &pb.CapResponse{
		ImgId:   id,
		ImgUrl:  url,
		ImgData: content.Bytes(),
	}, nil
}

func (e *Captcha) Validate(ctx context.Context, req *pb.VRequest) (*pb.VResponse, error) {
	res := captcha.VerifyString(req.Id, req.Value)
	return &pb.VResponse{
		Res: res,
	}, nil
}
