package cfg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/PatrikOlin/butler-burton/util"
	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name       string `yaml:"name"`
	Color      string `yaml:"color"`
	WebhookURL string `yaml:"webhook_url"`
	Report     Report `yaml:"report"`
}

type Report struct {
	Path               string `yaml:"path"`
	Update             bool   `yaml:"update"`
	StartingRow        int    `yaml:"starting_row"`
	StartingDayOfMonth int    `yaml:"starting_day_of_month"`
	CheckinCol         string `yaml:"checkin_col"`
	CheckoutCol        string `yaml:"checkout_col"`
	LunchCol           string `yaml:"lunch_col"`
}

var Cfg Config

func InitConfig() {
	dir := os.Getenv("HOME") + "/.config/butlerburton/"
	util.MakeDirectoryIfNotExists(dir)

	configFile := path.Join(dir, "config.yml")

	err := cleanenv.ReadConfig(configFile, &Cfg)
	if err != nil {
		createDefaultConfig(configFile)
	}
}

func createDefaultConfig(path string) {
	fmt.Println("failed to read config file, creating config with default values")
	Cfg = Config{
		Name:       "Burton",
		Color:      "#46D9FF",
		WebhookURL: "",
		Report: Report{
			Path:               "/home/olin/.butlerburton/",
			Update:             false,
			StartingRow:        12,
			StartingDayOfMonth: 16,
			CheckinCol:         "C",
			CheckoutCol:        "D",
			LunchCol:           "F",
		},
	}

	bytes, err := yaml.Marshal(Cfg)
	if err != nil {
		fmt.Println("failed to marshal default config values")
	}

	e := ioutil.WriteFile(path, bytes, 0644)
	if e != nil {
		fmt.Println("failed to create default config file")
		panic(e)
	}
}
