package main

import (
	"flag"
	"fmt"
)

var parameter = flag.String("fruit", "", "Please input the fruit you like")

func main() {
	flag.Parse()
	var fruit = map[string]func(){
		"apple":      func() { fmt.Println("🍎 <=> Apple") },
		"strawberry": func() { fmt.Println("🍓 <=> Strawberry") },
		"onion":      func() { fmt.Println("🧅 <=> Onion") },
	}
	if fn, ok := fruit[*parameter]; ok {
		fn()
	} else {
		fmt.Println("Fruit not found")
	}
}
