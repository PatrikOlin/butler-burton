package cfg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/PatrikOlin/butler-burton/util"
	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name         string `yaml:"name"`
	Color        string `yaml:"color"`
	WebhookURL   string `yaml:"webhook_url"`
	Notifcations bool   `yaml:"notifications"`
	VabMsg       string `yaml:"vab_msg"`
	Report       Report `yaml:"report"`
}

type Report struct {
	Path               string `yaml:"path"`
	Update             bool   `yaml:"update"`
	StartingRow        int    `yaml:"starting_row"`
	StartingDayOfMonth int    `yaml:"starting_day_of_month"`
	CheckinCol         string `yaml:"checkin_col"`
	CheckoutCol        string `yaml:"checkout_col"`
	LunchCol           string `yaml:"lunch_col"`
	BLLunchCol         string `yaml:"bl_lunch_col"`
	OvertimeCol        string `yaml:"overtime_col"`
	FlexInCol          string `yaml:"flex_in_col"`
	VabCol             string `yaml:"vab_col"`
}

var Cfg Config

func InitConfig() {
	fpath := GetConfigPath()

	err := cleanenv.ReadConfig(fpath, &Cfg)
	if err != nil {
		createDefaultConfig(fpath)
	}
}

func GetConfigPath() string {
	dir := os.Getenv("HOME") + "/.config/butlerburton/"
	util.MakeDirectoryIfNotExists(dir)

	fpath := path.Join(dir, "config.yml")

	return fpath
}

func ReloadConfig() {
	fpath := GetConfigPath()
	err := cleanenv.ReadConfig(fpath, &Cfg)
	if err != nil {
		log.Fatalln("Could not read config file, time to panic!")
	}

}

func createDefaultConfig(path string) {
	fmt.Println("failed to read config file, creating config with default values")
	Cfg = Config{
		Name:         "Burton",
		Color:        "#46D9FF",
		WebhookURL:   "",
		Notifcations: true,
		VabMsg:       "Jag vabbar idag, försök hålla skutan flytande så är jag tillbaka imorgon",
		Report: Report{
			Path:               "/home/olin/.butlerburton/",
			Update:             false,
			StartingRow:        12,
			StartingDayOfMonth: 16,
			CheckinCol:         "C",
			CheckoutCol:        "D",
			LunchCol:           "F",
			BLLunchCol:         "I",
			OvertimeCol:        "R",
			FlexInCol:          "V",
			VabCol:             "L",
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
