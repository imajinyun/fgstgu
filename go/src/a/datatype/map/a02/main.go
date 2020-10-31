package main

import "fmt"

func main() {
	countries := make(map[string]string)
	countries["China"] = "Beijing"
	countries["Japan"] = "Tokyo"
	countries["Italy"] = "Rome"
	countries["Korea"] = "Pyongyang"
	countries["India"] = "New Delhi"

	capital, ok := countries["Singapore"]
	fmt.Printf("%q\n", capital)
	fmt.Printf("ok type: %T, value: %v\n", ok, ok)

	if ok {
		fmt.Printf("Capital: %v\n", capital)
	} else {
		fmt.Println("Not found")
	}

	if capital, ok := countries["China"]; ok {
		fmt.Printf("Capital: %v\n", capital)
	} else {
		fmt.Println("Not found")
	}
}
