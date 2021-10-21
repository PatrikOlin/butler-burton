package cmd

import (
	"strings"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/util"
	"github.com/PatrikOlin/butler-burton/xlsx"
)

func DownloadReport() error {
	monthFolderReplacer := GetMonthFolderReplacer()
	monthFileReplacer := GetMonthFileReplacer()

	m := util.GetMonth()
	monthFolder := monthFolderReplacer.Replace(m)
	monthFile := monthFileReplacer.Replace(m)
	department := "Dev"

	name := strings.Replace(cfg.Cfg.Name, " ", "_", -1)

	util.DownloadReport(monthFolder, monthFile, department, name)
	xlsx.GetTransferableStock()

	return nil
}
