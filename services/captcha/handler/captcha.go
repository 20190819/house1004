package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dchest/captcha"
	pb "house1004/services/captcha/proto"
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
