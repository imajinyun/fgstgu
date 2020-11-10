package main

import (
	"bytes"
	"fmt"
	"math/big"
)

var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefgihjklmnpqrstuvwxyz")

func main() {
	arr := []byte("小楼昨夜又东风，故国不堪回首月明中。")
	encoded := Base58Encode(arr)
	decoded := Base58Decode(encoded)
	fmt.Println("Base58 encoded:\n", string(encoded))
	fmt.Println("Base58 decoded:\n", string(decoded))
}

// Base58Encode func.
func Base58Encode(input []byte) []byte {
	var res []byte
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(int64(len(base58Alphabets)))
	zero := big.NewInt(0)
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		res = append(res, base58Alphabets[mod.Int64()])
	}
	if input[0] == 0x00 {
		res = append(res, base58Alphabets[0])
	}
	ReverseBytes(res)
	return res
}

// Base58Decode func.
func Base58Decode(input []byte) []byte {
	res := big.NewInt(0)
	for _, b := range input {
		index := bytes.IndexByte(base58Alphabets, b)
		res.Mul(res, big.NewInt(58))
		res.Add(res, big.NewInt(int64(index)))
	}
	decoded := res.Bytes()
	if input[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded
}

// ReverseBytes func.
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
