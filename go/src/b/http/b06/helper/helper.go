package helper

import (
	"fmt"
	"os"
)

// CheckError func.
func CheckError(err error) {
	defer func() {
		if res, ok := recover().(error); ok {
			fmt.Println("✖️ An exception has occurred: ", res.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}

// IsFileExist func.
func IsFileExist(file string) bool {
	fileInfo, err := os.Stat(file)
	fmt.Println(fileInfo.Name())
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
