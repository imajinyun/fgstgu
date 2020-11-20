package main

import "fmt"

type someStruct struct{}

func (ss *someStruct) String() string {
	return "Hello"
}

// GetStringer func.
func GetStringer() fmt.Stringer {
	var ss *someStruct = nil
	return ss
}

// GetStringer2 func.
func GetStringer2() fmt.Stringer {
	var ss *someStruct = nil
	if ss == nil {
		return nil
	}
	return ss
}

func main() {
	if GetStringer() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}

	if GetStringer2() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}
}
