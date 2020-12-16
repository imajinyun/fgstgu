package main

import (
	"fgstgu/go/src/a/struct/a07/calendar"
	"fmt"
	"log"
)

func main() {
	d1 := calendar.Date{}
	fmt.Println(d1)

	d2 := calendar.Date{}
	if err := d2.SetYear(2020); err != nil {
		log.Fatalln(err)
	}
	if err := d2.SetMonth(12); err != nil {
		log.Fatalln(err)
	}
	if err := d2.SetDay(13); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(d2)
	fmt.Printf("Year: %d, Month: %d, Day: %d", d2.Year(), d2.Month(), d2.Day())
}
