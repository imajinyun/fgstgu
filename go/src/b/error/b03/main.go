package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("Custom error #1")
	fmt.Printf("%T, %v, %v\n", err1, err1.Error(), err1)

	err2 := fmt.Errorf("The type of error is #%d", 500)
	fmt.Printf("%T, %v, %v\n", err2, err2.Error(), err2)

	res, err := checkUserAge(-20)
	if err != nil {
		fmt.Printf("%T, %v, %T, %v\n", res, res, err, err.Error())
	} else {
		fmt.Println(res)
	}
}

func checkUserAge(age int) (string, error) {
	if age < 0 {
		err := fmt.Errorf("Your age cannot be negative")
		return "", err
	}
	return fmt.Sprintf("Your input age is %d\n", age), nil
}
