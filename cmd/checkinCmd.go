package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
	"github.com/PatrikOlin/butler-burton/xlsx"
)

func Checkin(opts util.Options) error {
	d := (15 * time.Minute)
	rounded := time.Now().Local().Round(d)
	checkinUnix := time.Now().Unix()
	db.Store.Put("checkinUnix", checkinUnix)
	db.Store.Put("checkinRounded", rounded)

	de := time.Unix(checkinUnix, 0).Local().Format("15:04:05")
	dr := rounded.Format("15:04:05")
	checkinMsg := fmt.Sprintf("Ok, checked in at %s (%s)\n", de, dr)
	fmt.Println(checkinMsg)

	if opts.Loud {
		util.SendTeamsMessage(
			fmt.Sprintf("%s checkar in", cfg.Cfg.Name),
			"Incheckad från "+string(de),
			cfg.Cfg.Color, cfg.Cfg.WebhookURL)
	}

	if cfg.Cfg.TimeSheet.Update {
		xlsx.SetCheckInCellValue(rounded, opts.Verbose)
	}

	if cfg.Cfg.Notifications {
		util.Notify("Checking in \n", checkinMsg)
	}

	menu, err := util.GetTodaysLunchMenu()
	if err != nil {
		return err
	} else {
		prettyPrintMenu(menu)
	}

	return nil
}

func VabCheckin(opts util.Options) error {
	if opts.Loud {
		util.SendTeamsMessage(
			fmt.Sprintf("%s vabbar", cfg.Cfg.Name),
			cfg.Cfg.VabMsg,
			cfg.Cfg.Color,
			cfg.Cfg.WebhookURL,
		)
	}

	if cfg.Cfg.TimeSheet.Update {
		xlsx.SetVabCheckin()
	}

	if cfg.Cfg.Notifications {
		util.Notify("Reporting vab \n", "")
	}

	return nil
}

func prettyPrintMenu(menu []util.AbbrMenuItem) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Rätt", "Kategori"})
	for _, item := range menu {
		t.AppendRow([]interface{}{item.Number, item.Name, item.Category})
	}
	t.AppendFooter(table.Row{"", "Du har väl inte glömt att beställa käk?", ""})

	t.Render()
}
