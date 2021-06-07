package cfg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

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
	EmployeeID       string `yaml:"employee_id"`
	Path             string `yaml:"path"`
	Update           bool   `yaml:"update"`
	CheckinCol       string `yaml:"checkin_col"`
	CheckoutCol      string `yaml:"checkout_col"`
	LunchCol         string `yaml:"lunch_col"`
	BLLunchCol       string `yaml:"bl_lunch_col"`
	OvertimeCol      string `yaml:"overtime_col"`
	VabCol           string `yaml:"vab_col"`
	AFKCol           string `yaml:"afk_col"`
	EmployeeIDCoords string `yaml:"employee_id_coords"`
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
	makeDirectoryIfNotExists(dir)

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
		Name:         os.Getenv("USER"),
		Color:        "#46D9FF",
		WebhookURL:   "",
		Notifcations: true,
		VabMsg:       "Jag vabbar idag, försök hålla skutan flytande så är jag tillbaka imorgon",
		Report: Report{
			EmployeeID:       "0000",
			Path:             os.Getenv("HOME") + "/.butlerburton/",
			Update:           false,
			EmployeeIDCoords: "C2",
			CheckinCol:       "C",
			CheckoutCol:      "D",
			LunchCol:         "F",
			BLLunchCol:       "I",
			OvertimeCol:      "R",
			VabCol:           "L",
			AFKCol:           "G",
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

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}
