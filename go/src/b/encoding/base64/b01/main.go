package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "大漠孤烟直，长河落日圆"
	encoded := Base64EncodeString(str)
	decoded := Base64DecodeString(encoded)
	fmt.Println("Base64 encoded:", encoded)
	fmt.Println("Base64 decoded:", decoded)
}

// Base64EncodeString func.
func Base64EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64DecodeString func.
func Base64DecodeString(str string) string {
	res, _ := base64.StdEncoding.DecodeString(str)
	return string(res)
}
