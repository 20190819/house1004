package main

import (
	"github.com/gin-gonic/gin"
	"house1004/web/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteString("go web 项目开始啦 ！")
	})
	r.GET("/api/v1/captcha", controllers.GetCaptcha)

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
