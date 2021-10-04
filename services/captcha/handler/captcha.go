package handler

import (
	"context"
	pb "house1004/services/captcha/proto"
)

type Captcha struct{}

func (e *Captcha) Call(ctx context.Context, req *pb.CapRequest) (*pb.CapResponse, error) {

	return &pb.CapResponse{
		Data: []byte("this is captcha service"),
	}, nil
}
