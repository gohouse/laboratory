package main

import (
	"time"
	"fmt"
)

const DATE_FORMAT = "2006-01-02"
const DATETIME_FORMAT = "2006-01-02 15:04:05"

func main() {
	res := GetDate()
	fmt.Println(res)
	// ===============
	res2,_ := time.Parse(DATETIME_FORMAT, res.YesterdayStart)
	y,m,d := res2.Date()
	day := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	dayEnd := time.Date(y, m, d, 23, 59, 59, 999, time.Local)
	fmt.Println(day.Format(DATE_FORMAT+" 00:00:00"), dayEnd.Format(DATETIME_FORMAT))
}

type DateTime struct {
	LastMonthStart string
	LastMonthEnd string
	ThisMonthStart string
	TodayStart string
	YesterdayStart string
	Now string
}

func tmp() {
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	fmt.Println(today.Format(DATE_FORMAT+" 00:00:00"))
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(thisMonth.Format(DATE_FORMAT))
	start := thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	timeRange := fmt.Sprintf("%s~%s", start, end)
	fmt.Println(timeRange)
}

func GetDate() DateTime {
	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	yeaterday := today.AddDate(0, 0, -1)
	return DateTime{
		LastMonthStart: thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT),
		LastMonthEnd: thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT),
		ThisMonthStart: thisMonth.Format(DATE_FORMAT),
		TodayStart: today.Format(DATE_FORMAT+" 00:00:00"),
		YesterdayStart: yeaterday.Format(DATE_FORMAT+" 00:00:00"),
		Now: now.Format(DATETIME_FORMAT),
	}
}
