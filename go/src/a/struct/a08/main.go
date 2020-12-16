package main

import (
	"fgstgu/go/src/a/struct/a08/calendar"
	"fmt"
	"log"
)

func main() {
	e1 := calendar.Event{}
	if err := e1.Date.SetYear(2020); err != nil {
		log.Fatalln(err)
	}
	if err := e1.Date.SetMonth(12); err != nil {
		log.Fatalln(err)
	}
	if err := e1.Date.SetDay(13); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Year: %v, Month: %v, Day: %v\n", e1.Year(), e1.Month(), e1.Day())

	e2 := calendar.Event{}
	if err := e2.SetYear(2020); err != nil {
		log.Fatalln(err)
	}
	if err := e2.SetMonth(12); err != nil {
		log.Fatalln(err)
	}
	if err := e2.SetDay(13); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Year: %v, Month: %v, Day: %v\n", e2.Year(), e2.Month(), e2.Day())

	e3 := calendar.Event{}
	if err := e3.SetTitle("🎉 Hello World, Hello Go!"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(e3.Title())
}
