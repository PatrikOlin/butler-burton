package util

import (
	"time"
)

func GetMonth() string {
	var month string
	reportBreakPoint := 15

	d := time.Now().Day()
	if d > reportBreakPoint {
		month = time.Now().AddDate(0, 1, 0).Local().Month().String()
	} else {
		month = time.Now().Local().Month().String()
	}

	return month
}
