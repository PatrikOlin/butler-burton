package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/PatrikOlin/skvs"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/urfave/cli/v2"
)

var store *skvs.KVStore

type TeamsMessage struct {
	Title      string `json:"title"`
	ThemeColor string `json:"themeColor"`
	Text       string `json:"text"`
}

type Config struct {
	Name       string `yaml:"name"`
	WebhookURL string `yaml:"webhook_url"`
}

var cfg Config

func init() {
	initDB()
	initConfig()
}

func initConfig() {
	dir := os.Getenv("HOME") + "/.config/butlerburton/"
	makeDirectoryIfNotExists(dir)

	configFile := path.Join(dir, "config.yml")

	err := cleanenv.ReadConfig(configFile, &cfg)
	if err != nil {
		fmt.Println("failed to read config file, using default values")
		cfg = Config{
			Name:       "Burton",
			WebhookURL: "",
		}
	}
}

func initDB() {
	dir := os.Getenv("HOME") + "/.butlerburton/"
	makeDirectoryIfNotExists(dir)

	dbfile := path.Join(dir, "data.db")

	var err error
	store, err = skvs.Open(dbfile)
	if err != nil {
		log.Fatal(err)
	}
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
					d := (15 * time.Minute)
					rounded := time.Now().Local().Round(d)
					checkinUnix := time.Now().Unix()
					store.Put("checkinUnix", checkinUnix)
					store.Put("checkinRounded", rounded)

					de := time.Unix(checkinUnix, 0).Local().Format("15:04:05")
					dr := rounded.Format("15:04:05")
					fmt.Printf("Ok, checked in at %s (%s)\n", de, dr)
					sendTeamsMessage(
						fmt.Sprintf("%s checkar in", cfg.Name),
						"Incheckad från "+string(de))
					return nil
				},
			},
			{
				Name:    "check out",
				Aliases: []string{"co"},
				Usage:   "trigger check out sequence",
				Action: func(c *cli.Context) error {
					var valUnix int64
					if err := store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
						fmt.Println("not found")
						return err
					} else if err != nil {
						log.Fatal(err)
						return err
					}

					var valRound time.Time
					if err := store.Get("checkinRounded", &valRound); err == skvs.ErrNotFound {
						fmt.Println("not found")
						return err
					} else if err != nil {
						log.Fatal(err)
						return err
					} else {
						fmt.Println("Ok, checking out.")
						fmt.Printf("Time spent checked in: %s\n", calculateTimeCheckedIn(valUnix))
						de := time.Unix(valUnix, 0).Local().Format("15:04:05")
						dr := valRound.Local().Format("15:04:05")
						fmt.Printf("You checked in at: %s (%s)\n", de, dr)
						sendTeamsMessage(
							fmt.Sprintf("%s checkar ut", cfg.Name),
							"Utcheckad från "+string(time.Now().Format("15:04:05")))
					}
					return nil
				},
			},
			{
				Name:    "check time",
				Aliases: []string{"ct"},
				Usage:   "get time spent checked in",
				Action: func(c *cli.Context) error {
					var valUnix int64
					if err := store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
						fmt.Println("not found")
						return err
					} else if err != nil {
						log.Fatal(err)
						return err
					} else {
						fmt.Printf("Time checked in %s\n", calculateTimeCheckedIn(valUnix))
					}
					return nil

				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func calculateTimeCheckedIn(checkin int64) time.Duration {
	t1 := time.Unix(checkin, 0)
	t2 := time.Since(t1)

	d := (1000 * time.Millisecond)
	trunc := t2.Truncate(d)
	return trunc
}

func sendTeamsMessage(title, msg string) error {
	tBody, _ := json.Marshal(TeamsMessage{
		Title:      title,
		ThemeColor: "#ba7016",
		Text:       msg,
	})

	req, err := http.NewRequest(http.MethodPost, cfg.WebhookURL, bytes.NewBuffer(tBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Teams")
	}

	return nil
}
