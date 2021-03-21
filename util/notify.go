package util

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func Notify(header, body string) {
	err := beeep.Notify(header, body, "")
	if err != nil {
		fmt.Println(err)
	}
}
