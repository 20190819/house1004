package helpers

import (
	exceptions "house1004/exceptions"
	"strconv"
)

func String2Int(str string) int {
	i, err := strconv.Atoi(str)
	exceptions.Fatal(err)
	return i
}
