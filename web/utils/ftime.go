package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Ftime time.Time

const DEFAULT_TIME_FORMAT = "2006-01-02 15:04:05"

func (t Ftime) MarshalJSON() ([]byte, error) {
	tStr := time.Now().Format(DEFAULT_TIME_FORMAT) // 设置格式
	// 注意 json 字符串风格要求
	return []byte(fmt.Sprintf("\"%v\"", tStr)), nil
}

func (t Ftime) Value() (driver.Value, error) {
	// Ftime 转换成 time.Time 类型
	return time.Now().Format(DEFAULT_TIME_FORMAT), nil
}

func (t *Ftime) Scan(v interface{}) error {
	return nil
}
