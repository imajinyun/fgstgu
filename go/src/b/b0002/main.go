package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe("127.0.0.1:8080", nil))
}
