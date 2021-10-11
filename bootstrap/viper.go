package bootstrap

import (
	"fmt"
	exceptions "house1004/exceptions"

	"github.com/spf13/viper"
)

func LoadEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env") //  env 类型
	viper.AddConfigPath("./../../")
	if err := viper.ReadInConfig(); err != nil {
		exceptions.Fatal(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(".env 配置文件未找到")
			return
		}
	}
}
