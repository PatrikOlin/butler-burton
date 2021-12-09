package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/koltyakov/gosip"
	strategy "github.com/koltyakov/gosip-sandbox/strategies/azurecert"
	"github.com/koltyakov/gosip/api"
)

type LunchMenuItem []struct {
	Metadata struct {
		ID   string `json:"id"`
		URI  string `json:"uri"`
		Etag string `json:"etag"`
		Type string `json:"type"`
	} `json:"__metadata"`
	Day                  string    `json:"Day"`
	MenuX0020Item        int       `json:"Menu_x0020_Item"`
	ItemX0020Name        string    `json:"Item_x0020_Name"`
	ItemX0020Description string    `json:"Item_x0020_Description"`
	Week                 string    `json:"Week"`
	DagNr                string    `json:"DagNr"`
	Created              time.Time `json:"Created"`
}

func DownloadBaseReport(name, monthFolder, monthFile string) string {
	y := time.Now().Format("2006")
	y2 := time.Now().Format("06")
	fileRelativeURL := "Tidrapporter/" + y + "/" + monthFolder + "/TIRP_Original_" + monthFile + "-" + y2 + ".xlsx"

	fileName := "TIRP_" + name + "_" + monthFile + "-" + y2 + ".xlsx"

	getFile(fileRelativeURL, fileName)
	return fileName
}

func DownloadReport(monthFolder, monthFile, department, name string) error {
	y := time.Now().Format("2006")
	y2 := time.Now().Format("06")
	fileRelativeURL := "Tidrapporter/" + y + "/" + monthFolder + "/" +
		department + "/TIRP_" + name + "_" + monthFile + "-" + y2 + ".xlsx"

	fileName := "TIRP_" + name + "_" + monthFile + "-" + y2 + ".xlsx"

	getFile(fileRelativeURL, fileName)
	return nil
}

func UploadReport(monthFolder, department, filePath string) error {
	sp := auth()
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	p := strings.Split(filePath, "/")
	fileName := p[len(p)-1]

	y := time.Now().Format("2006")
	folder := sp.Web().GetFolder(fmt.Sprintf("Tidrapporter/%s/%s/%s/", y, monthFolder, department))

	_, err = folder.Files().Add(fileName, contents, true)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GetTodaysLunchMenu() error {
	url := "http://bltv01.blinfo.se:4300/lunch/" + strconv.Itoa(getWeek())

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	str := string(body)

	return nil
}

func filterTodaysMeals() {

}

func getFile(fileRelURL, fileName string) {
	sp := auth()
	data, err := sp.Web().GetFile(fileRelURL).Download()
	if err != nil {
		log.Fatal(err)
	}

	filePath := os.Getenv("HOME") + "/.butlerburton/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("unable to create a file: %v\n", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("unable to write to file: %v\n", err)
	}

	file.Sync()
}

func auth() *api.SP {
	authCnfg := &strategy.AuthCnfg{}
	configPath := os.Getenv("HOME") + "/.butlerburton/private.json"
	if err := authCnfg.ReadConfig(configPath); err != nil {
		log.Fatalf("unable to get config: %v", err)
	}

	client := &gosip.SPClient{AuthCnfg: authCnfg}
	sp := api.NewSP(client)

	return sp
}

func getWeek() int {
	t := time.Now().UTC()
	_, week := t.ISOWeek()

	return week
}

func getWeekDay() time.Weekday {
	return time.Now().Weekday()
}
