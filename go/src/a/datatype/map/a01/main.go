package main

import "fmt"

func main() {
	var cities = map[string]string{
		"Shangxi": "Xi'an",
		"Fujian":  "Fuzhou",
		"Hebei":   "Shijiazhuang",
		"Gansu":   "Lanzhou",
	}
	fmt.Printf("%v\n", cities)

	others := map[string]float64{"Python": 5, "Go": 99.99, "Java": 19.2, "C++": 88.23}
	fmt.Println(others)

	countries := make(map[string]string)
	countries["China"] = "Beijing"
	countries["Japan"] = "Tokyo"
	countries["Italy"] = "Rome"
	countries["Korea"] = "Pyongyang"

	for k, v := range countries {
		fmt.Printf("Country: %v, Capital: %v\n", k, v)
	}

	for _, v := range countries {
		fmt.Printf("Country capital: %v\n", v)
	}

	for k := range countries {
		fmt.Printf("Country: %v\n", k)
	}
}
