package routes

import "github.com/gin-gonic/gin"

func RegisterWebRoutes(r *gin.Engine){
	// 加载静态文件
	r.Static("/home", "./web/views")
}