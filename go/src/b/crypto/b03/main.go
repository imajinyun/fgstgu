package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"log"
	"os"
)

// 生成 RSA 密钥文件
func main() {
	var bits int
	flag.IntVar(&bits, "b", 1024, "Key length, default is 1024 bits")
	if err := GenRsaKey(bits); err != nil {
		log.Fatalln("Key file generation failed")
	}
}

// GenRsaKey func.
func GenRsaKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	bytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{Type: "Private Key", Bytes: bytes}
	file, err := os.Create("/tmp/private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey
	bytes, err = x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{Type: "Public Key", Bytes: bytes}
	file, err = os.Create("/tmp/public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
