package handler

import (
	"context"
	"house1004/exceptions"
	"house1004/services/user/models/user"
	pb "house1004/services/user/proto"
)

type User struct {
}

func (e User) Register(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	var usr *user.User
	usr=&user.User{
		Name: req.Phone,
		Mobile: req.Phone,
		Password: req.Password,
	}
	err := usr.Create()
	if err != nil {
		exceptions.Fatal(err)
	}
	return &pb.Response{}, nil
}
