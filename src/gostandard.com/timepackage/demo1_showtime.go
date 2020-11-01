package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	Year := t.Year()
	Month := t.Month()
	Day := t.Day()

	fmt.Printf("Today is %d/%d/%d\n", Month, Day, Year)

	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format("Monday, January 2 in the year 2006"))

	startDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC)
	fmt.Println(startDate)
	fmt.Println(startDate.Format(layoutISO))
	fmt.Println(startDate.Format(layoutUS))
}
