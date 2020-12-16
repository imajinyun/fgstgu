package main

import (
	"fgstgu/go/src/b/http/b06/helper"
	"fmt"
	"net/http"
)

// HTMLPrefix is html prefix.
const HTMLPrefix = "/test"

// PublicDir is root path.
const PublicDir = http.Dir("/tmp/")

func main() {
	addr := "127.0.0.1:8081"
	mux := http.NewServeMux()
	fmt.Printf("Please click http://%s to access it.\n", addr)
	mux.HandleFunc(HTMLPrefix, handler)
	err := http.ListenAndServe(addr, mux)
	helper.CheckError(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[1:len(HTMLPrefix)])
	file := string(PublicDir) + r.URL.Path[1:len(HTMLPrefix)] + ".txt"
	fmt.Println(file)
	if ok := helper.IsFileExist(file); ok {
		w.Write([]byte("🍎 Very good!"))
		return
	}
	http.ServeFile(w, r, file)
}
