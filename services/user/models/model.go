package models

import (
	exceptions "house1004/exceptions"
	"house1004/services"
	"house1004/services/user/models/user"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Migration 迁移文件
func Migration() {
	err := services.DB.AutoMigrate(&user.User{})
	exceptions.Fatal(err)
}
