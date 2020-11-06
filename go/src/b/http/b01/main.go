package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer()
}

func fileServer() {
	root, port := http.Dir("./src/"), ":8080"
	log.Printf(" File server root path is %s\n", root)
	log.Printf(" Please click http://127.0.0.1%s to access it\n", port)
	http.ListenAndServe(port, http.FileServer(root))
}
