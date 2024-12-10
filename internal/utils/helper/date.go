package helper

import (
	"fmt"
	"time"
)

var layout = "2006-01-02 15:04"

var Hari = map[string]string{
	"Sunday":    "Minggu",
	"Monday":    "Senin",
	"Tuesday":   "Selasa",
	"Wednesday": "Rabu",
	"Thursday":  "Kamis",
	"Friday":    "Jumat",
	"Saturday":  "Sabtu",
}

var Bulan = map[time.Month]string{
	time.January:   "Januari",
	time.February:  "Februari",
	time.March:     "Maret",
	time.April:     "April",
	time.May:       "Mei",
	time.June:      "Juni",
	time.July:      "Juli",
	time.August:    "Agustus",
	time.September: "September",
	time.October:   "Oktober",
	time.November:  "November",
	time.December:  "Desember",
}

func ParseDate(date string) (jam string, hari string, tanggal string) {
	t, _ := time.Parse(layout, date)
	jam = fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
	hari = Hari[t.Weekday().String()]
	tanggal = fmt.Sprintf("%d %s %d", t.Day(), Bulan[t.Month()], t.Year())

	return jam, hari, tanggal
}

func ParseUTC(date string) string {
	t, _ := time.Parse(layout, date)

	return t.Format("2006/01/02 15:04:05 UTC")
}
