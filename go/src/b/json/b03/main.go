package main

import (
	"encoding/json"
	"fmt"
)

// User struct.
type User struct {
	Name    string `json:"_name"`
	Age     int    `json:"_age"`
	Gender  uint   `json:"-"`
	Address string
}

var user = []User{
	{"Jack", 70, 0, "Beijing"},
	{"Jhon", 59, 1, "Shanghai"},
}

func main() {
	for _, v := range user {
		arr, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(arr))
	}
}
