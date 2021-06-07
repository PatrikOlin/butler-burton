package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/koltyakov/gosip"
	strategy "github.com/koltyakov/gosip-sandbox/strategies/azurecert"
	"github.com/koltyakov/gosip/api"
)

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
	configPath := os.Getenv("HOME") + "/.config/butlerburton/private.json"
	if err := authCnfg.ReadConfig(configPath); err != nil {
		log.Fatalf("unable to get config: %v", err)
	}

	client := &gosip.SPClient{AuthCnfg: authCnfg}
	sp := api.NewSP(client)

	return sp
}
