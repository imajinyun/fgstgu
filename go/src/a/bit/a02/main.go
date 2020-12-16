package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("%b&%b == %b\n", 0, 0, 0&0)
	log.Printf("%b&%b == %b\n", 0, 1, 0&1)
	log.Printf("%b&%b == %b\n", 1, 1, 1&1)

	log.Println()
	log.Printf("%b&%b == %b\n", 0, 0, 0|0)
	log.Printf("%b&%b == %b\n", 0, 1, 0|1)
	log.Printf("%b&%b == %b\n", 1, 1, 1|1)

	log.Println()
	log.Printf("%08b\n", 1)
	log.Printf("%08b\n", 2)
	log.Printf("%08b\n", 4)
	log.Printf("%08b\n", 8)
	log.Printf("%08b\n", 16)
	log.Printf("%08b\n", 32)
	log.Printf("%08b\n", 64)
	log.Printf("%08b\n", 128)

	log.Println()
	log.Printf("%016b\n", os.O_RDONLY)
	log.Printf("%016b\n", os.O_WRONLY)
	log.Printf("%016b\n", os.O_RDWR)
	log.Printf("%016b\n", os.O_APPEND)
	log.Printf("%016b\n", os.O_SYNC)
	log.Printf("%016b\n", os.O_CREATE)
	log.Printf("%016b\n", os.O_TRUNC)
	log.Printf("%016b\n", os.O_EXCL)

	log.Println()
	log.Println(os.FileMode(0700))
	log.Println(os.FileMode(0644))
	log.Println(os.FileMode(0755))
	log.Println(os.FileMode(0777))

	log.Println()
	log.Printf("%09b\n", 0007)
	log.Printf("%09b\n", 0070)
	log.Printf("%09b\n", 0700)

	log.Println()
	log.Printf("%09b %s\n", 0000, os.FileMode(0000))
	log.Printf("%09b %s\n", 0070, os.FileMode(0070))
	log.Printf("%09b %s\n", 0700, os.FileMode(0700))
	log.Printf("%09b %s\n", 0644, os.FileMode(0644))
}
