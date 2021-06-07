package cmd

import (
	"strings"
)

func GetMonthFolderReplacer() *strings.Replacer {
	r := strings.NewReplacer(
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
