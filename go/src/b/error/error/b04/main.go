package main

import (
	"fmt"
	"time"
)

// MyError struct.
type MyError struct {
	startTime time.Time
	message   string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Time: %v, Error: %v\n", e.startTime, e.message)
}

func getArea(width, height float64) (float64, error) {
	err := ""
	if width < 0 && height < 0 {
		err = fmt.Sprintf("width=%v, height=%v\n", width, height)
	} else {
		if width < 0 {
			err = fmt.Sprintf("weight=%v\n", width)
		}
		if height < 0 {
			err = fmt.Sprintf("height=%v\n", height)
		}
	}
	if err != "" {
		return 0, MyError{time.Now(), err}
	}
	return width * height, nil
}

func main() {
	res, err := getArea(-10, -10)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
