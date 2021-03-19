package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/PatrikOlin/skvs"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
	"github.com/PatrikOlin/butler-burton/xlsx"
)

func Checkout(blOpt bool, verbose bool) error {
	var valUnix int64
	if err := db.Store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	}
	var valRound time.Time
	if err := db.Store.Get("checkinRounded", &valRound); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	} else {
		fmt.Println("Ok, checking out.")
		checkedInMsg := fmt.Sprintf("Time spent checked in: %s\n", CalculateTimeCheckedIn(valUnix))
		fmt.Println(checkedInMsg)

		de := time.Unix(valUnix, 0).Local().Format("15:04:05")
		dr := valRound.Local().Format("15:04:05")

		d := (15 * time.Minute)
		roundedNow := time.Now().Local().Round(d)

		checkedInDurMsg := fmt.Sprintf("You checked in at: %s (%s)\n", de, dr)
		fmt.Println(checkedInDurMsg)
		util.SendTeamsMessage(
			fmt.Sprintf("%s checkar ut", cfg.Cfg.Name),
			"Utcheckad från "+string(time.Now().Format("15:04:05")),
			cfg.Cfg.Color,
			cfg.Cfg.WebhookURL)

		if cfg.Cfg.Report.Update {
			xlsx.SetCheckOutCellValue(roundedNow, blOpt, verbose)
		}

		if cfg.Cfg.Notifcations {
			n := fmt.Sprintf("%s%s \n", checkedInMsg, checkedInDurMsg)
			util.Notify("Checking out \n", n)
		}

	}
	return nil
}

func CalculateTimeCheckedIn(checkin int64) time.Duration {
	t1 := time.Unix(checkin, 0)
	t2 := time.Since(t1)

	d := (1000 * time.Millisecond)
	trunc := t2.Truncate(d)
	return trunc
}
