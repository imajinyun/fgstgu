package main

import (
	"log"
	"os"

	"github.com/zieckey/etcdsync"
)

func main() {
	m, err := etcdsync.New("/mylock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Println("etcdsync.New failed:", err)
		return
	}
	m.SetDebugLogger(os.Stdout)
	if err = m.Lock(); err != nil {
		log.Println("etcdsync.Lock failed:", err)
	} else {
		log.Println("etcdsync.Lock successfully")
	}
	log.Println("get the lock. do something here...")
	if err = m.Unlock(); err != nil {
		log.Println("etcdsync.Unlock failed:", err)
	} else {
		log.Println("etcdsync.Unlock successfully")
	}
}
