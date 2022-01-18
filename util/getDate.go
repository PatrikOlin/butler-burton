package util

import (
	"time"
)

func GetMonth() string {
	var month string
	month = time.Now().Local().Month().String()

	return month
}
