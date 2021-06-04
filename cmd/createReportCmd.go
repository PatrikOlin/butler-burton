package cmd

import (
	"strings"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/util"
	"github.com/PatrikOlin/butler-burton/xlsx"
)

func CreateNewReport() error {
	monthFolderReplacer := strings.NewReplacer(
		"January", "01. januari",
		"February", "02. februari",
		"March", "03. mars",
		"April", "04. april",
		"May", "05. maj",
		"June", "06. juni",
		"July", "07. juli",
		"August", "08. augusti",
		"September", "09. september",
		"October", "10. oktober",
		"November", "11. november",
		"December", "12. december",
	)
	monthFileReplacer := strings.NewReplacer(
		"January", "Jan",
		"February", "Feb",
		"March", "Mars",
		"April", "April",
		"May", "Maj",
		"June", "Juni",
		"July", "Juli",
		"August", "Aug",
		"September", "Sep",
		"October", "Okt",
		"November", "Nov",
		"December", "Dec",
	)
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
