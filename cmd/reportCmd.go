package cmd

import (
	"fmt"
	"log"

	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/skvs"
)

func SetReportName(name string) error {
	db.Store.Put("reportName", name)
	fmt.Printf("Gotcha, set %s as report name\n", name)
	return nil
}

func GetReportName() error {
	var rn string
	if err := db.Store.Get("reportName", &rn); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	} else {
		fmt.Printf("Current report name is %s\n", rn)
	}
	return nil
}
