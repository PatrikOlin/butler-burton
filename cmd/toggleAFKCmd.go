package cmd

import (
	"fmt"
	"time"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
	"github.com/PatrikOlin/butler-burton/xlsx"
	"github.com/PatrikOlin/skvs"
)

func ToggleAFK(reason string, opts util.Options) error {
	var isAFK bool
	if err := db.Store.Get("isAFK", &isAFK); err == skvs.ErrNotFound {
		setAFK(reason, opts)
	} else if err != nil {
		return err
	} else if !isAFK {
		setAFK(reason, opts)
	} else {
		removeAFK(reason, opts)
	}

	return nil
}

func setAFK(reason string, opts util.Options) {
	t := time.Now().Unix()
	db.Store.Put("isAFK", true)
	db.Store.Put("AFKStartUnix", t)

	var isAFK bool
	var tu int64
	db.Store.Get("isAFK", &isAFK)
	db.Store.Get("AFKStartUnix", &tu)

	if cfg.Cfg.Notifications {
		util.Notify("Checkar ut som AFK \n", time.Now().Format("15:04:05"))
	}

	var msg string
	if reason != "" {
		msg = reason
	} else {
		msg = "Checkar ut, återkommer senare idag"
	}

	if !opts.Silent {
		util.SendTeamsMessage(
			fmt.Sprintf("%s checkar ut en stund", cfg.Cfg.Name),
			msg,
			cfg.Cfg.Color,
			cfg.Cfg.WebhookURL,
		)
	}
}

func removeAFK(reason string, opts util.Options) {
	var t1 int64

	db.Store.Put("isAFK", false)
	if err := db.Store.Get("AFKStartUnix", &t1); err == skvs.ErrNotFound {
		fmt.Println("AFK start time not found in db")
	} else if err != nil {
		fmt.Println(err)
	} else {
		t2 := time.Now().Unix()
		t3 := calculateDuration(t1, t2)

		d := (15 * time.Minute)
		dur := t3.Round(d)

		db.Store.Put("AFKDuration", dur)
		db.Store.Delete("AFKStartUnix")

		var msg string
		if reason != "" {
			msg = reason
		} else {
			msg = " "
		}

		if !opts.Silent {
			util.SendTeamsMessage(
				fmt.Sprintf("Äntligen är %s tillbaka!", cfg.Cfg.Name),
				msg,
				cfg.Cfg.Color,
				cfg.Cfg.WebhookURL,
			)
		}

		if cfg.Cfg.Notifications {
			util.Notify("Checkar in igen \n", time.Now().Format("15:04:05"))
		}

		if cfg.Cfg.Report.Update {
			xlsx.SetAFKCellValue(fmtDuration(dur))
		}
	}
}

func calculateDuration(unix1, unix2 int64) time.Duration {
	t1 := time.Unix(unix1, 0)
	t2 := time.Unix(unix2, 0)

	dur := t2.Sub(t1)

	return dur
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}
