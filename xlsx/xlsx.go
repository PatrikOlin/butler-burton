package xlsx

import (
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/util"
)

func SetCheckInCellValue(ciTime time.Time, verbose bool) {
	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)
	cellCoords := cfg.ColCfg.CheckinCol + getRowNumber(f, sheet)

	p, err := util.GetFilePath()
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
	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)

	row := getRowNumber(f, sheet)
	vabCoords := cfg.ColCfg.VabCol + row

	f.SetCellValue(sheet, vabCoords, "08:00")

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetCheckOutCellValue(coTime time.Time, ot string, catering, verbose bool) {
	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)

	row := getRowNumber(f, sheet)
	cellCoords := cfg.ColCfg.CheckoutCol + row
	lunchCoords := cfg.ColCfg.LunchCol + row

	f.SetCellValue(sheet, cellCoords, coTime.Format("15:04"))
	f.SetCellValue(sheet, lunchCoords, "01:00")
	p, err := util.GetFilePath()
	if err != nil {
		fmt.Println(err)
		return
	}

	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", coTime.Format("15:04"), cellCoords, p)
		fmt.Printf("Writing %s to cell %s in %s\n", "01:00", lunchCoords, p)
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

func SetAFKCellValue(AFKDuration string) {
	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)

	row := getRowNumber(f, sheet)
	AFKCoords := cfg.ColCfg.AFKCol + row

	f.SetCellValue(sheet, AFKCoords, AFKDuration)

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetEmployeeID() {
	f, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	i := f.GetActiveSheetIndex()
	sheet := f.GetSheetName(i)

	eID := cfg.Cfg.Report.EmployeeID
	eIDCoords := cfg.ColCfg.EmployeeIDCoords

	f.SetCellValue(sheet, eIDCoords, eID)

	err = f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func openFile() (*excelize.File, error) {
	p, err := util.GetFilePath()
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
	blLunchCoords := cfg.ColCfg.BLLunchCol + row
	p, err := util.GetFilePath()
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Duration(1 * time.Hour)
	f.SetCellFormula(sheet, blLunchCoords, "")
	f.SetCellValue(sheet, blLunchCoords, t)
	if verbose == true {
		fmt.Printf("Writing %s to cell %s in %s\n", t, blLunchCoords, p)
	}
}

func setOvertime(ot string, f *excelize.File, sheet, row string, verbose bool) {
	overtimeCoords := cfg.ColCfg.OvertimeCol + row

	p, err := util.GetFilePath()
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

func getRowNumber(f *excelize.File, sheet string) string {
	var rowNum string
	t1 := time.Now()
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
	}

	for i, _ := range rows {
		coord := "A" + strconv.Itoa(i)
		val, _ := f.GetCellValue(sheet, coord)
		t2, _ := time.Parse("01-02-06", val)
		t2 = t2.AddDate(0, 0, 1)
		if t1.Month() == t2.Month() && t1.Day() == t2.Day() {
			rowNum = strconv.Itoa(i)
			break
		}
	}

	return rowNum
}

// func getPath() (string, error) {
// 	var rn string
// 	if err := db.Store.Get("reportFilename", &rn); err == skvs.ErrNotFound {
// 		log.Fatal("not found")
// 		return "", err
// 	} else if err != nil {
// 		log.Fatal(err)
// 		return "", err
// 	} else {
// 		return path.Join(cfg.Cfg.Report.Path, rn), nil
// 	}
// }
