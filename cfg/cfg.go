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
	Name          string    `yaml:"name"`
	Color         string    `yaml:"color"`
	WebhookURL    string    `yaml:"webhook_url"`
	Notifications bool      `yaml:"notifications"`
	VabMsg        string    `yaml:"vab_msg"`
	TimeSheet     TimeSheet `yaml:"time_sheet"`
}

type TimeSheet struct {
	EmployeeID string `yaml:"employee_id"`
	Path       string `yaml:"path"`
	Update     bool   `yaml:"update"`
}

type ColumnConfig struct {
	CheckinCol                    string
	CheckoutCol                   string
	LunchCol                      string
	BLLunchCol                    string
	OvertimeCol                   string
	VabCol                        string
	AFKCol                        string
	ExerciseCol                   string
	EmployeeIDCoords              string
	TransferredPositiveFlexCoords string
	TransferredNegativeFlexCoords string
	TransferredCompTimeCoords     string
	OutgoingFlexCoords            string
	OutgoingCompTimeCoords        string
	OutgoingATFCoords             string
}

var Cfg Config
var ColCfg ColumnConfig

func InitConfig() {
	ColCfg = createColumnConfig()
	fpath := GetConfigPath()

	err := cleanenv.ReadConfig(fpath, &Cfg)
	if err != nil {
		fmt.Println("failed to read config file, creating config with default values")
		createDefaultConfig(fpath)
	}
}

func GetConfigPath() string {
	dir := os.Getenv("HOME") + "/.config/butlerburton/"
	err := makeDirectoryIfNotExists(dir)
	if err != nil {
		fmt.Println(err)
	}

	fpath := path.Join(dir, "config.yml")

	return fpath
}

func ReloadConfig() {
	fpath := GetConfigPath()
	err := cleanenv.ReadConfig(fpath, &Cfg)
	if err != nil {
		log.Fatalln("Could not read config file.")
	}

}

func createDefaultConfig(path string) {
	Cfg = Config{
		Name:          os.Getenv("USER"),
		Color:         "#46D9FF",
		WebhookURL:    "",
		Notifications: true,
		VabMsg:        "Jag vabbar idag, försök hålla skutan flytande så är jag tillbaka imorgon",
		TimeSheet: TimeSheet{
			EmployeeID: "0000",
			Path:       os.Getenv("HOME") + "/.butlerburton/",
			Update:     false,
		},
	}

	bytes, err := yaml.Marshal(Cfg)
	if err != nil {
		fmt.Println("failed to marshal default config values")
	}

	e := ioutil.WriteFile(path, bytes, 0644)
	if e != nil {
		fmt.Println("failed to create default config file")
		fmt.Println(e)
		panic(e)
	}
}

func createColumnConfig() ColumnConfig {
	return ColumnConfig{
		CheckinCol:                    "C",
		CheckoutCol:                   "D",
		LunchCol:                      "F",
		BLLunchCol:                    "I",
		OvertimeCol:                   "R",
		VabCol:                        "L",
		AFKCol:                        "G",
		ExerciseCol:                   "J",
		EmployeeIDCoords:              "C2",
		TransferredPositiveFlexCoords: "S2",
		TransferredNegativeFlexCoords: "S3",
		TransferredCompTimeCoords:     "S4",
		OutgoingFlexCoords:            "T2",
		OutgoingCompTimeCoords:        "T4",
		OutgoingATFCoords:             "T5",
	}
}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModeDir|0755)
	}
	return nil
}
