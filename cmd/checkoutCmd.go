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

func Checkout(blOpt, overtime, verbose bool) error {
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
		tci := CalculateTimeCheckedIn(valUnix)
		fmt.Println("Ok, checking out.")
		checkedInMsg := fmt.Sprintf("Time spent checked in: %s\n", tci)
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

		var ot string

		if cfg.Cfg.Report.Update && overtime {
			ot = calculateOvertime(tci)
		}

		if cfg.Cfg.Report.Update && ot != "" {
			xlsx.SetCheckOutCellValue(roundedNow, ot, blOpt, verbose)
		} else if cfg.Cfg.Report.Update {
			xlsx.SetCheckOutCellValue(roundedNow, "", blOpt, verbose)
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

func calculateOvertime(tci time.Duration) string {
	ot := tci - (9 * time.Hour)
	r := (15 * time.Minute)
	ot = ot.Round(r)
	hh := ot / time.Hour
	ot -= hh * time.Hour
	mm := ot / time.Minute
	return fmt.Sprintf("%02d:%02d", hh, mm)
}
