package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	arr := []byte{'H', 'e', 'l', 'l', 'o', ',', ' ', 'G', 'o', '!'}
	fmt.Println(string(arr))

	str := BytesToHexString(arr)
	fmt.Println(str)

	str = ReverseHexString(str)
	arr, _ = HexStringToBytes(str)
	fmt.Printf("%x\n", arr)
	ReverseBytes(arr)
	fmt.Println(string(arr))
}

// BytesToHexString func.
func BytesToHexString(arr []byte) string {
	return hex.EncodeToString(arr)
}

// HexStringToBytes func.
func HexStringToBytes(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

// ReverseHexString func.
func ReverseHexString(str string) string {
	arr, _ := hex.DecodeString(str)
	ReverseBytes(arr)
	return hex.EncodeToString(arr)
}

// ReverseBytes func.
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
