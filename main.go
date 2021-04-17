package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/cmd"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
)

var version string

func init() {
	db.InitDB()
	cfg.InitConfig()
}

func main() {
	opts := util.Options{
		Verbose:  false,
		Catered:  false,
		Overtime: false,
		Vab:      false,
	}

	app := &cli.App{
		Name:    "Butler Burton",
		Usage:   "Your personal butler",
		Version: version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Value:       false,
				Usage:       "Turn on verbose mode",
				Destination: &opts.Verbose,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "check in",
				Aliases: []string{"ci"},
				Usage:   "trigger check in sequence",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "vab",
						Aliases:     []string{"v"},
						Value:       false,
						Usage:       "Check in as absent as result of vab",
						Destination: &opts.Vab,
					},
				},
				Action: func(c *cli.Context) error {
					if opts.Vab {
						return cmd.VabCheckin()
					}
					return cmd.Checkin(opts)
				},
			},
			{
				Name:    "check out",
				Aliases: []string{"co"},
				Usage:   "trigger check out sequence",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "catered",
						Aliases:     []string{"c"},
						Value:       false,
						Usage:       "Check BL-lunch field in report for todays shift",
						Destination: &opts.Catered,
					},
					&cli.BoolFlag{
						Name:        "overtime",
						Aliases:     []string{"o"},
						Value:       false,
						Usage:       "Write overtime to overtime column",
						Destination: &opts.Overtime,
					},
				},
				Action: func(c *cli.Context) error {
					return cmd.Checkout(opts)
				},
			},
			{
				Name:    "check time",
				Aliases: []string{"ct"},
				Usage:   "get time spent checked in",
				Action: func(c *cli.Context) error {
					return cmd.CheckTime()
				},
			},
			{
				Name:     "report",
				Aliases:  []string{"r"},
				Usage:    "report commands",
				Category: "report",
				Subcommands: []*cli.Command{
					{
						Name:     "set",
						Aliases:  []string{"s"},
						Usage:    "set new report filename",
						Category: "report",
						Action: func(c *cli.Context) error {
							return cmd.SetReportFilename(c.Args().First())
						},
					},
					{
						Name:     "get",
						Aliases:  []string{"g"},
						Usage:    "get current report filename",
						Category: "report",
						Action: func(c *cli.Context) error {
							return cmd.GetReportFilename()
						},
					},
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "edit config-file",
				Action: func(c *cli.Context) error {
					return cmd.EditConfig()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
