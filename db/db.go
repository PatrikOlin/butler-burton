package db

import (
	"log"
	"os"
	"path"

	"github.com/PatrikOlin/skvs"

	"github.com/PatrikOlin/butler-burton/util"
)

var Store *skvs.KVStore

func InitDB() {
	dir := os.Getenv("HOME") + "/.butlerburton/"
	util.MakeDirectoryIfNotExists(dir)

	dbfile := path.Join(dir, "data.db")

	var err error
	Store, err = skvs.Open(dbfile)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCheckinTime() (int64, error) {
	var valUnix int64
	if err := Store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
		log.Println("not found")
		return -1, err
	} else if err != nil {
		log.Fatal(err)
		return -1, err
	}

	return valUnix, nil
}
