package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/PatrikOlin/skvs"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
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
		checkInTimeMsg := fmt.Sprintf("You checked in at %s\n", time.Unix(valUnix, 0).Local().Format("15:04:05"))
		timeCheckedInMsg := fmt.Sprintf("Time checked in %s\n", CalculateTimeCheckedIn(valUnix))
		fmt.Printf(checkInTimeMsg)
		fmt.Printf(timeCheckedInMsg)

		if cfg.Cfg.Notifications {
			n := fmt.Sprintf("%s%s \n", checkInTimeMsg, timeCheckedInMsg)
			util.Notify("Checked in duration \n", n)
		}
	}
	return nil
}
