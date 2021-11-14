package exceptions

import (
	"errors"
	"log"
)

var (
	ErrParams = errors.New("参数异常")
	ErrNotFound = errors.New("数据未找到")
)

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
