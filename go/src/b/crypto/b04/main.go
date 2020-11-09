package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// RSA 加密与解密
func main() {
	var (
		err                   error
		decrypted             string
		privateKey, publicKey []byte
	)
	flag.StringVar(&decrypted, "d", "", "Encrypted data")
	flag.Parse()
	publicKey, err = ioutil.ReadFile("/tmp/public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("/tmp/private.pem")
	if err != nil {
		os.Exit(-1)
	}
	fmt.Printf("%T, %T\n", publicKey, privateKey)

	str := "人之初，性本善；性相近，习相远。"
	fmt.Println("Original:", str)

	encrypter, _ := RsaEncryptString(str, publicKey)
	fmt.Println("Encrypted:", encrypter)

	decrypter, _ := RsaDecryptString(encrypter, privateKey)
	fmt.Println("Decrypted: ", decrypter)
}

// RsaEncryptString func.
func RsaEncryptString(str string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("Public key error")
	}
	parser, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	public := parser.(*rsa.PublicKey)
	cipherArr, err := rsa.EncryptPKCS1v15(rand.Reader, public, []byte(str))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherArr), nil
}

// RsaDecryptString func.
func RsaDecryptString(cipher string, privateKey []byte) (string, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("Private key error")
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	cipherArr, _ := base64.StdEncoding.DecodeString(cipher)
	sourceArr, err := rsa.DecryptPKCS1v15(rand.Reader, private, cipherArr)
	if err != nil {
		return "", err
	}
	return string(sourceArr), nil

}
