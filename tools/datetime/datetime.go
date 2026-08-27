package datetime

import (
	"fmt"
	"time"
)

var (
	day = map[string]string{
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
	}

	month = map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}
)

// IndonesiaDate to convert time to indonesia format (10 Januari 2022)
func IndonesiaDate(t time.Time) string {
	months := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	return fmt.Sprintf("%d %s %d", t.Day(), months[t.Month()-1], t.Year())
}

// GenFullIndonesianDate to generate full indonesian date, ex : Senin, 25 Oktober 2024
func GenFullIndonesianDate(t time.Time) string {
	dayFr := t.Format("Monday")
	dayDate := t.Format("2")
	monthFr := t.Format("January")
	yearDate := t.Format("2006")

	return fmt.Sprintf("%s, %s %s %s", day[dayFr], dayDate, month[monthFr], yearDate)
}

// GenIndonesianDateTime to generate indonesian date time, ex : 25 Oktober 2024, 13:30
func GenIndonesianDateTime(t time.Time) string {
	dayDate := t.Format("2")
	monthFr := t.Format("January")
	yearDate := t.Format("2006")
	hour24Fr := t.Format("15")
	minuteFr := t.Format("04")

	return fmt.Sprintf("%s %s %s, %s:%s", dayDate, month[monthFr], yearDate, hour24Fr, minuteFr)
}
