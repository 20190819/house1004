package user

import (
	"house1004/bootstrap"
	"house1004/exceptions"
	"house1004/services/user/models"
)

type User struct {
	models.BaseModel
	Name     string `gorm:"size:32;unique"`  //用户名
	Password string `gorm:"size:128" `       //用户密码加密的
	Mobile   string `gorm:"size:11;unique" ` //手机号
	RealName string `gorm:"size:32" `        //真实姓名  实名认证
	IdCard   string `gorm:"size:20" `        //身份证号  实名认证
	Avatar   string `gorm:"size:256" `       //用户头像路径
}

// Migration 迁移文件
func Migration() {
	err := bootstrap.DB.AutoMigrate(&User{})
	exceptions.Fatal(err)
}
