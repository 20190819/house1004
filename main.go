package main

import (
	exceptions "house1004/exceptions"
	"house1004/web/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.RegisterWebRoutes(r)
	routes.RegisterApiRoutes(r)

	err := r.Run(":8888")
	exceptions.Fatal(err)
}
