package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/PatrikOlin/skvs"

	"github.com/PatrikOlin/butler-burton/db"
)

func CheckTime() error {
	var valUnix int64
	if err := db.Store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	} else {
		fmt.Printf("You checked in at %s\n", time.Unix(valUnix, 0).Local().Format("15:04:05"))
		fmt.Printf("Time checked in %s\n", CalculateTimeCheckedIn(valUnix))

	}
	return nil
}
