package handler

import (
	"context"
	"fmt"
	"house1004/bootstrap"
	"house1004/exceptions"
	"house1004/services/user/models/user"
	pb "house1004/services/user/proto"
)

type User struct {
}

func (e User) Register(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	var usr *user.User
	usr = &user.User{
		Name:     req.Phone,
		Mobile:   req.Phone,
		Password: req.Password,
	}
	err := usr.Create()
	if err != nil {
		exceptions.Fatal(err)
	}
	return &pb.Response{}, nil
}

func (e User) Login(ctx context.Context, req *pb.LRequest) (*pb.LResponse, error) {
	var usr *user.User
	result := bootstrap.DB.Where(req).First(&usr)
	fmt.Println("login usr>>>", result.Error)
	if result.RowsAffected == 0 {
		// 未找到
		return &pb.LResponse{},exceptions.ErrNotFound
	} else {
		// 返回用信息
		return &pb.LResponse{}, nil
	}

}
