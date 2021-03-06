package xlsx

import (
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/PatrikOlin/butler-burton/cfg"
)

func SetCheckInCellValue(ciTime time.Time) {
	f, err := excelize.OpenFile(getPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	cellCoords := getCellCoords(cfg.Cfg.Report.CheckinCol)
	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Writing %s to cell %s in %s\n", ciTime, cellCoords, getPath())
	f.SetCellValue(sheet, cellCoords, ciTime.Format("15:04"))

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetCheckOutCellValue(coTime time.Time) {
	lunchDuration, err := time.Parse("15:04", "01:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := excelize.OpenFile(getPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	cellCoords := getCellCoords(cfg.Cfg.Report.CheckoutCol)
	lunchCoords := getCellCoords(cfg.Cfg.Report.LunchCol)
	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)

	fmt.Printf("Writing %s to cell %s in %s\n", coTime.Format("15:04"), cellCoords, getPath())
	f.SetCellValue(sheet, cellCoords, coTime.Format("15:04"))
	fmt.Printf("Writing %s to cell %s in %s\n", lunchDuration.Format("15:04"), lunchCoords, getPath())
	f.SetCellValue(sheet, lunchCoords, lunchDuration.Format("15:04"))

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getCellCoords(col string) string {
	offset := cfg.Cfg.Report.StartingRow
	t := time.Now().Local()
	currDay := t.Day()
	var daysSinceStart int

	var t1 time.Time
	if currDay < 15 {
		t1 = time.Now().Local().AddDate(0, -1, (cfg.Cfg.Report.StartingDayOfMonth - currDay))
	} else {
		t1 = time.Now().Local().AddDate(0, 0, (cfg.Cfg.Report.StartingDayOfMonth - currDay))
	}
	t2 := time.Now().Local()
	daysSinceStart = int(t2.Sub(t1).Hours() / 24)

	return col + strconv.Itoa(daysSinceStart+offset)
}

func getPath() string {
	return path.Join(cfg.Cfg.Report.Path, "burtontest.xlsx")
}
