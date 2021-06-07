package cmd

import (
	"github.com/PatrikOlin/butler-burton/util"
)

func UploadReport() error {
	monthFolderReplacer := GetMonthFolderReplacer()

	m := util.GetMonth()
	monthFolder := monthFolderReplacer.Replace(m)
	department := "Dev"

	fp, err := util.GetFilePath()
	if err != nil {
		return err
	}

	util.UploadReport(monthFolder, department, fp)

	return nil
}
