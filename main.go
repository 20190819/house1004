package main

import (
	bootstrap2 "house1004/bootstrap"
	exceptions "house1004/exceptions"
	"house1004/web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap2.LoadEnv()

	bootstrap2.ConnectDB()
	bootstrap2.Migration()

	r := gin.Default()
	routes.RegisterWebRoutes(r)
	routes.RegisterApiRoutes(r)

	err := r.Run(":8888")
	exceptions.Fatal(err)
}
