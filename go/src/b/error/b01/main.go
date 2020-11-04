package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	res := math.Sqrt(-100)
	fmt.Println(res)

	res, err := Sqrt(-100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = Divide(100, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	file, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}
}

// Sqrt func.
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		err := "Negative numbers have no square roots."
		return 0, errors.New(err)
	}
	return math.Sqrt(f), nil
}

// Divide func.
func Divide(divide float64, divisor float64) (float64, error) {
	if divisor == 0 {
		err := "Divisor cannot be zero."
		return 0, errors.New(err)
	}
	return divide / divisor, nil
}
