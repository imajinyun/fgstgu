package main

import (
	"fmt"
	"os"
)

func main() {
	filename1 := "/tmp/abc.txt"
	file1, err1 := os.Open(filename1)
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		fmt.Printf("%v: %s open successfully!\n", file1, filename1)
	}

	filename2 := "/tmp/test.txt"
	file2, err2 := os.OpenFile(filename2, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Printf("%v: %s open successfully!\n", file2, filename2)
	}
	file1.Close()
	file2.Close()

	filename3 := "/tmp/test.txt"
	err3 := os.Remove(filename3)
	if err3 != nil {
		fmt.Println(err3.Error())
	} else {
		fmt.Printf("%s remove successfully!\n", filename3)
	}

	filename4 := "/tmp/test"
	err4 := os.RemoveAll(filename4)
	if err4 != nil {
		fmt.Println(err4.Error())
	} else {
		fmt.Printf("%s remove successfully!\n", filename4)
	}
}
