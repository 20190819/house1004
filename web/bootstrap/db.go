package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"house1004/web/exceptions"
	"house1004/web/models/user"
	"time"
)

var DB *gorm.DB

func ConnectDB(){
	dns:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_DATABASE"),
		)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: viper.GetString("DB_CONNECTION"),
		DSN: dns,
	}), &gorm.Config{})
	exceptions.Handler(err)
	DB=db
	pool()
}

// 设置连接池
func pool(){
	sqlDB, err := DB.DB()
	exceptions.Handler(err)

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// Migration 迁移文件
func Migration(){
	err:=DB.AutoMigrate(&user.User{})
	exceptions.Handler(err)
}

