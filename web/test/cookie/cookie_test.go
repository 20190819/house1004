package cookie

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

type User struct {
	Id int	`json:"id"`
	Name string	`json:"name"`
	Age int8	`json:"age"`
}

func TestCookie(t *testing.T){
	router:=gin.Default()
	user:=&User{1,"lisi",18}
	uStr,_:=json.Marshal(&user)
	fmt.Println(string(uStr))
	router.GET("/", func(ctx *gin.Context) {
		ctx.SetCookie("user",string(uStr),300,"","",false,false)
		ctx.Writer.WriteString("this is a test cookie")
	})

	router.Run(":7000")
}