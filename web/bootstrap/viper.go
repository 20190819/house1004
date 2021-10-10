package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"house1004/web/exceptions"
)

func LoadEnv(){
	viper.SetConfigName(".env")
	viper.SetConfigType("env")  //  env 类型
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		exceptions.Handler(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(".env 配置文件未找到")
			return
		}
	}
}