package util

import (
	"time"
)

func GetMonth() string {
	return time.Now().Local().Month().String()
}
