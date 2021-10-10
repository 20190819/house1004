package main

import (
	"github.com/gin-gonic/gin"
	"house1004/web/bootstrap"
	"house1004/web/exceptions"
	"house1004/web/routes"
)

func main() {
	bootstrap.LoadEnv()

	bootstrap.ConnectDB()
	bootstrap.Migration()

	r := gin.Default()
	routes.RegisterWebRoutes(r)
	routes.RegisterApiRoutes(r)

	err := r.Run(":8888")
	exceptions.Handler(err)
}
