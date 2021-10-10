package exceptions

import (
	"errors"
	"log"
)

var (
	ErrParams = errors.New("参数异常")
)

func Handler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
