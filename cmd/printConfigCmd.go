package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/PatrikOlin/butler-burton/cfg"
)

func PrintConfig() error {
	configFile := cfg.GetConfigPath()

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("Failed to read file: %s", err)
	}
	fmt.Printf("CONFIG FILE:\n   %s\n", configFile)
	fmt.Printf("\nSIZE:\n   %d bytes\n", len(data))
	fmt.Printf("\nDATA:\n%s", data)
	return err
}
