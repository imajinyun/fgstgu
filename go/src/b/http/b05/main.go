package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := "127.0.0.1:8081"
	http.HandleFunc("/index", indexHandler)
	fmt.Printf("Please click http://%s/index to access\n", addr)
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🔥 The back end received the request...")
	w.Write([]byte("😂 It's a default page!"))
}
