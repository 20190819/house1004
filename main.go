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
	var cc controllers.CaptchaController
	r.GET("/api/v1/captcha/cd", cc.GetCaptchaCd)
	r.GET("/api/v1/captcha/:id/img", cc.GetCaptchaImg)
	r.POST("/api/v1/captcha/:id/verify", cc.Verify)

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
