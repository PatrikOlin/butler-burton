package util

import (
	"log"
	"path"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/skvs"
)

func GetFilePath() (string, error) {
	var rn string
	if err := db.Store.Get("reportFilename", &rn); err == skvs.ErrNotFound {
		log.Fatal("not found")
		return "", err
	} else if err != nil {
		log.Fatal(err)
		return "", err
	} else {
		return path.Join(cfg.Cfg.Report.Path, rn), nil
	}
}
