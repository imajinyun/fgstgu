package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

// 椭圆曲线数字签名算法 - ECDSA
func main() {
	privateKey, publicKey := NewKeyPair()
	message := sha256.Sum256([]byte("Go Language"))
	r, s, _ := ecdsa.Sign(rand.Reader, &privateKey, message[:])
	strSignR, strSignS := fmt.Sprintf("%x", r), fmt.Sprintf("%x", s)
	fmt.Printf("r dec: %v\ns dec: %v\n", r, s)
	fmt.Printf("r hex: %v\ns hex: %v\n", strSignR, strSignS)
	signatureDer := MakeSignatureDerString(strSignR, strSignS)
	fmt.Println("Signature format:\n", signatureDer)
	res := VerifySign(publicKey, message[:], r, s)
	fmt.Println("Verify result:", res)
	res = VerifySignature(publicKey, message[:], strSignR, strSignS)
	fmt.Println("Signature result:", res)
}

// NewKeyPair func.
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	publicKey := append(private.X.Bytes(), private.Y.Bytes()...)
	return *private, publicKey
}

// VerifySign func.
func VerifySign(publicKey, message []byte, r, s *big.Int) bool {
	pklen := len(publicKey)
	x, y := big.Int{}, big.Int{}
	x.SetBytes(publicKey[:(pklen / 2)])
	y.SetBytes(publicKey[(pklen / 2):])
	rawPublicKey := ecdsa.PublicKey{elliptic.P256(), &x, &y}
	res := ecdsa.Verify(&rawPublicKey, message, r, s)
	return res
}

// VerifySignature func.
func VerifySignature(publicKey, message []byte, r, s string) bool {
	pklen := len(publicKey)
	x, y := big.Int{}, big.Int{}
	x.SetBytes(publicKey[:(pklen / 2)])
	y.SetBytes(publicKey[(pklen / 2):])
	rawPublicKey := ecdsa.PublicKey{elliptic.P256(), &x, &y}
	rint, sint := big.Int{}, big.Int{}
	rByte, _ := hex.DecodeString(r)
	sByte, _ := hex.DecodeString(s)
	rint.SetBytes(rByte)
	sint.SetBytes(sByte)
	res := ecdsa.Verify(&rawPublicKey, message, &rint, &sint)
	return res
}

// MakeSignatureDerString func.
func MakeSignatureDerString(r, s string) string {
	lenSignR, lenSignS := len(r)/2, len(s)/2
	lenSequence := lenSignR + lenSignS + 4
	strLenSignR := DecToHex(int64(lenSignR))
	strLenSignS := DecToHex(int64(lenSignS))
	strLenSequence := DecToHex(int64(lenSequence))
	str := "30" + strLenSequence
	str = str + "02" + strLenSignR + r
	str = str + "02" + strLenSignS + s
	str = str + "01"
	return str
}

// DecToHex func.
func DecToHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}
