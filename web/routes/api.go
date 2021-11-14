package routes

import (
	"github.com/gin-gonic/gin"
	"house1004/web/controllers"
)

func RegisterApiRoutes(r *gin.Engine){

	api:=r.Group("/api/v1")
	{
		var cc controllers.CaptchaController
		captcha:=api.Group("/captcha")
		captcha.GET("/cd",cc.GetCaptchaCd)
		captcha.GET("/:id/img",cc.GetCaptchaImg)
		captcha.POST("/:id/verify",cc.Verify)

		var uc controllers.UserController
		user:=api.Group("/user")
		user.POST("/register",uc.HandlerRegister)
		user.POST("/login",uc.HandlerLogin)
	}
}