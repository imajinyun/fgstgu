package main

import "fmt"

func main() {
Looper:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue Looper
			}
		}
	}
}
