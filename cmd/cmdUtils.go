package cmd

import (
	"strings"
)

func GetMonthFolderReplacer() *strings.Replacer {
	r := strings.NewReplacer(
		"January", "01. Januari",
		"February", "02. Februari",
		"March", "03. Mars",
		"April", "04. April",
		"May", "05. Maj",
		"June", "06. Juni",
		"July", "07. Juli",
		"August", "08. Augusti",
		"September", "09. September",
		"October", "10. Oktober",
		"November", "11. November",
		"December", "12. December",
	)
	return r
}

func GetMonthFileReplacer() *strings.Replacer {
	r := strings.NewReplacer(
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
	return r
}
