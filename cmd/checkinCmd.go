package cmd

import (
	"fmt"
	"time"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
	"github.com/PatrikOlin/butler-burton/xlsx"
)

func Checkin(verbose bool) error {
	d := (15 * time.Minute)
	rounded := time.Now().Local().Round(d)
	checkinUnix := time.Now().Unix()
	db.Store.Put("checkinUnix", checkinUnix)
	db.Store.Put("checkinRounded", rounded)

	de := time.Unix(checkinUnix, 0).Local().Format("15:04:05")
	dr := rounded.Format("15:04:05")
	fmt.Printf("Ok, checked in at %s (%s)\n", de, dr)
	util.SendTeamsMessage(
		fmt.Sprintf("%s checkar in", cfg.Cfg.Name),
		"Incheckad från "+string(de),
		cfg.Cfg.Color, cfg.Cfg.WebhookURL)

	if cfg.Cfg.Report.Update {
		xlsx.SetCheckInCellValue(rounded, verbose)
	}
	return nil
}
