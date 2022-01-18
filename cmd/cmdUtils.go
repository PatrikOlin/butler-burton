package cmd

import (
	"strings"
)

func GetMonthFolderReplacer() *strings.Replacer {
	r := strings.NewReplacer(
		"January", "Januari",
		"February", "Februari",
		"March", "Mars",
		"April", "April",
		"May", "Maj",
		"June", "Juni",
		"July", "Juli",
		"August", "Augusti",
		"September", "September",
		"October", "Oktober",
		"November", "November",
		"December", "December",
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
