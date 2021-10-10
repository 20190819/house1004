package helpers

import (
	"house1004/web/exceptions"
	"strconv"
)

func String2Int(str string)int{
	i,err:=strconv.Atoi(str)
	exceptions.Handler(err)
	return i
}