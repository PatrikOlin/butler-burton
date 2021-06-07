package cmd

import (
	"strings"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/util"
	"github.com/PatrikOlin/butler-burton/xlsx"
)

func CreateNewReport() error {
	monthFolderReplacer := GetMonthFolderReplacer()
	monthFileReplacer := GetMonthFileReplacer()
	m := util.GetMonth()
	monthFolder := monthFolderReplacer.Replace(m)
	monthFile := monthFileReplacer.Replace(m)
	name := strings.Replace(cfg.Cfg.Name, " ", "_", -1)
	reportName := util.DownloadBaseReport(name, monthFolder, monthFile)

	err := SetReportFilename(reportName)
	if err != nil {
		return err
	}

	xlsx.SetEmployeeID()

	return nil
}
