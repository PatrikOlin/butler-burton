package xlsx

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/skvs"
)

func SetCheckInCellValue(ciTime time.Time, verbose bool) {
	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	cellCoords := cfg.Cfg.Report.CheckinCol + getRowNumber()
	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)
	if err != nil {
		fmt.Println(err)
	}

	p, err := getPath()
	if err != nil {
		fmt.Println(err)
		return
	}

	f.SetCellValue(sheet, cellCoords, ciTime.Format("15:04"))

	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", ciTime.Format("15:04"), cellCoords, p)
	}

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetVabCheckin() {
	row := getRowNumber()
	vabCoords := cfg.Cfg.Report.VabCol + row

	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)

	f.SetCellValue(sheet, vabCoords, "08:00")

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetCheckOutCellValue(coTime time.Time, ot string, catering, verbose bool) {
	lunchDuration, err := time.Parse("15:04", "01:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	row := getRowNumber()

	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	cellCoords := cfg.Cfg.Report.CheckoutCol + row
	lunchCoords := cfg.Cfg.Report.LunchCol + row
	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)

	f.SetCellValue(sheet, cellCoords, coTime.Format("15:04"))
	f.SetCellValue(sheet, lunchCoords, time.Duration(1*time.Hour))
	p, err := getPath()
	if err != nil {
		fmt.Println(err)
		return
	}

	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", coTime.Format("15:04"), cellCoords, p)
		fmt.Printf("Writing %s to cell %s in %s\n", lunchDuration.Format("15:04"), lunchCoords, p)
	}

	if catering == true {
		setCateredLunch(f, sheet, row, verbose)
	}

	if ot != "" {
		setOvertime(ot, f, sheet, row, verbose)
	}

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func openFile() (*excelize.File, error) {
	p, err := getPath()
	if err != nil {
		return nil, err
	}

	f, err := excelize.OpenFile(p)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return f, nil
}

func setCateredLunch(f *excelize.File, sheet, row string, verbose bool) {
	blLunchCoords := cfg.Cfg.Report.BLLunchCol + row
	str := "1"
	p, err := getPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellFormula(sheet, blLunchCoords, "")
	f.SetCellValue(sheet, blLunchCoords, str)
	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", str, blLunchCoords, p)
	}
}

func setOvertime(ot string, f *excelize.File, sheet, row string, verbose bool) {
	overtimeCoords := cfg.Cfg.Report.OvertimeCol + row

	p, err := getPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.SetCellValue(sheet, overtimeCoords, ot)
	if err != nil {
		fmt.Println(err)
		return
	}
	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", ot, overtimeCoords, p)
	}

}

func getRowNumber() string {
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

	return strconv.Itoa(daysSinceStart + offset)
}

func getPath() (string, error) {
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
