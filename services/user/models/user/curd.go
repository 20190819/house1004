package user

import (
	"house1004/bootstrap"
	"house1004/exceptions"
)

func (user *User)Create()error{
	err:=bootstrap.DB.Create(user).Error
	exceptions.Fatal(err)
	return err
}