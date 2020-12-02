/**
 * [HttpRouter](https://pkg.go.dev/github.com/julienschmidt/httprouter)
 * -> go get -u github.com/julienschmidt/httprouter
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index func.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "🎉 Welcome!\n")
}

// Hello func.
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "🐶 Hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
