package cmd

import (
	"fmt"
	"log"

	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/skvs"
)

func SetReportFilename(name string) error {
	db.Store.Put("reportFilename", name)
	fmt.Printf("Time sheet filename set to \n", name)
	return nil
}

func GetReportFilename() error {
	var rn string
	if err := db.Store.Get("reportFilename", &rn); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	} else {
		fmt.Printf("Current time sheet filename is %s\n", rn)
	}
	return nil
}
