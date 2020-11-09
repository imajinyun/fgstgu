package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

func main() {
	str := "Hello Go!"
	fmt.Println(str)

	a := Hash(str, "md5", false)
	fmt.Println(a)

	b := Hash(str, "sha1", false)
	fmt.Println(b)

	c := Hash(str, "sha256", false)
	fmt.Println(c)

	arr := SHA256Double(str, false)
	fmt.Println(string(arr))

	d := SHA256DoubleString(str, false)
	fmt.Println(d)
}

// Hash func.
func Hash(text string, hashType string, isHex bool) string {
	var instance hash.Hash
	switch hashType {
	case "md5":
		instance = md5.New()
	case "sha1":
		instance = sha1.New()
	case "sha256":
		instance = sha256.New()
	case "sha512":
		instance = sha512.New()
	}
	if isHex {
		arr, _ := hex.DecodeString(text)
		instance.Write(arr)
	} else {
		instance.Write([]byte(text))
	}
	cipher := instance.Sum(nil)
	return fmt.Sprintf("%x", cipher)
}

// SHA256Double func.
func SHA256Double(text string, isHex bool) []byte {
	instance := sha256.New()
	if isHex {
		arr, _ := hex.DecodeString(text)
		instance.Write(arr)
	} else {
		instance.Write([]byte(text))
	}
	cipher := instance.Sum(nil)
	instance.Reset()
	instance.Write(cipher)
	cipher = instance.Sum(nil)
	return cipher
}

// SHA256DoubleString func.
func SHA256DoubleString(text string, isHex bool) string {
	instance := sha256.New()
	if isHex {
		arr, _ := hex.DecodeString(text)
		instance.Write(arr)
	} else {
		instance.Write([]byte(text))
	}
	cipher := instance.Sum(nil)
	instance.Reset()
	instance.Write(cipher)
	cipher = instance.Sum(nil)
	return fmt.Sprintf("%x", cipher)
}
