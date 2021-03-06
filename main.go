package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/cmd"
	"github.com/PatrikOlin/butler-burton/db"
)

func init() {
	db.InitDB()
	cfg.InitConfig()
}

func main() {
	app := &cli.App{
		Name:  "Butler Burton",
		Usage: "Your personal butler",
		Commands: []*cli.Command{
			{
				Name:    "check in",
				Aliases: []string{"ci"},
				Usage:   "trigger check in sequence",
				Action: func(c *cli.Context) error {
					return cmd.Checkin()
				},
			},
			{
				Name:    "check out",
				Aliases: []string{"co"},
				Usage:   "trigger check out sequence",
				Action: func(c *cli.Context) error {
					return cmd.Checkout()
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
						Usage:    "set new report name",
						Category: "report",
						Action: func(c *cli.Context) error {
							return cmd.SetReportName(c.Args().First())
						},
					},
					{
						Name:     "get",
						Aliases:  []string{"g"},
						Usage:    "get current report name",
						Category: "report",
						Action: func(c *cli.Context) error {
							return cmd.GetReportName()
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
