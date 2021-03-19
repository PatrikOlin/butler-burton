package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/PatrikOlin/butler-burton/cfg"
)

func EditConfig() error {
	fpath := cfg.GetConfigPath()

	e := os.Getenv("EDITOR")
	cmd := exec.Command(e, fpath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("Error while editing. Error: %v\n", err)
	} else {
		fmt.Printf("Ok, saved. Lets test drive those changes!\n")
		cfg.ReloadConfig()
		fmt.Printf("Config reloaded\n")
		return err
	}
	return nil
}
