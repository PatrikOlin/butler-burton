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

	f.SetCellValue(sheet, cellCoords, fixTime(ciTime))
	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", ciTime.Format("15:04"), cellCoords, path)
	}

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetCheckOutCellValue(coTime time.Time, blOpt, overtime, verbose bool) {
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

	f.SetCellValue(sheet, cellCoords, fixTime(coTime))
	f.SetCellValue(sheet, lunchCoords, time.Duration(1*time.Hour))

	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", coTime.Format("15:04"), cellCoords, path)
		fmt.Printf("Writing %s to cell %s in %s\n", lunchDuration.Format("15:04"), lunchCoords, path)
	}

	if blOpt == true {
		setCateredLunch(f, sheet, row, verbose)
	}

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func openFile() (*excelize.File, error) {
	path, err := getPath()
	if err != nil {
		return nil, err
	}

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return f, nil
}

func setCateredLunch(f *excelize.File, sheet string, row string, verbose bool) {
	blLunchCoords := cfg.Cfg.Report.BLLunchCol + row
	str := "BL"
	path, err := getPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellFormula(sheet, blLunchCoords, "")
	f.SetCellValue(sheet, blLunchCoords, str)
	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", str, blLunchCoords, path)
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

func fixTime(t time.Time) time.Time {
	// workaround för konstig tidhantering av excelize, se https://github.com/360EntSecGroup-Skylar/excelize/issues/409
	t, _ = time.ParseInLocation("2006-01-02 15:04", t.Format("2006-01-02 15:04"), time.UTC)
	return t
}
