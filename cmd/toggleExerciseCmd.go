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

func ToggleExercise(msg string, opts util.Options) error {
	var isExercising bool
	if err := db.Store.Get("isExercising", &isExercising); err == skvs.ErrNotFound {
		setIsExercising(msg, opts)
	} else if err != nil {
		return err
	} else if !isExercising {
		setIsExercising(msg, opts)
	} else {
		removeIsExercising(opts)
	}
	return nil
}

func setIsExercising(msg string, opts util.Options) {
	t := time.Now().Unix()
	db.Store.Put("isExercising", true)
	db.Store.Put("exerciseStartUnix", t)
	m := fmt.Sprintf("Checked out for excercise at %s", time.Unix(t, 0).Local().Format("15:04:05"))
	fmt.Println(m)

	if cfg.Cfg.Notifications {
		util.Notify("Checkar ut för friskvård \n", time.Now().Format("15:04:05"))
	}

	var message string
	if msg != "" {
		message = msg
	} else {
		message = "Drar iväg och friskvårdar lite"
	}

	if !opts.Silent {
		util.SendTeamsMessage(
			fmt.Sprintf("%s checkar ut en stund", cfg.Cfg.Name),
			message,
			cfg.Cfg.Color,
			cfg.Cfg.WebhookURL,
		)
	}
}

func removeIsExercising(opts util.Options) {
	var t1 int64

	db.Store.Put("isExercising", false)
	if err := db.Store.Get("exerciseStartUnix", &t1); err == skvs.ErrNotFound {
		fmt.Println("Exercise start time not found in db")
	} else if err != nil {
		fmt.Println(err)
	} else {
		t2 := time.Now().Unix()
		t3 := calculateDurationLunchException(t1, t2)

		d := (15 * time.Minute)
		dur := t3.Round(d)

		db.Store.Put("exerciseDuration", dur)
		db.Store.Delete("exerciseStartUnix")
		msg := fmt.Sprintf("Checked in after excercising for %s", fmtDuration(dur))
		fmt.Println(msg)

		if !opts.Silent {
			util.SendTeamsMessage(
				fmt.Sprintf("Äntligen är %s tillbaka!", cfg.Cfg.Name),
				"",
				cfg.Cfg.Color,
				cfg.Cfg.WebhookURL,
			)
		}

		if cfg.Cfg.Notifications {
			util.Notify("Checkar in igen \n", time.Now().Format("15:04:05"))
		}

		if cfg.Cfg.Report.Update {
			xlsx.SetExerciseCellValue(fmtDuration(dur))
		}
	}
}

func calculateDurationLunchException(unix1, unix2 int64) time.Duration {
	t1 := time.Unix(unix1, 0)
	t2 := time.Unix(unix2, 0)

	lunchStart, err := time.Parse("15:04", "12:00")
	lunchEnd, err := time.Parse("15:04", "13:00")
	if err != nil {
		fmt.Println(err)
	}

	if t1.Local().Before(lunchStart) && t2.Local().After(lunchEnd) {
		t2.Add(-1 * time.Hour)
	}

	dur := t2.Sub(t1)

	return dur
}
