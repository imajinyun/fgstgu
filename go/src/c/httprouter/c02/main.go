package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Employee type.
type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employees map[string]*Employee

func init() {
	employees = map[string]*Employee{}
	employees["jack"] = &Employee{"E100001", "Jack", 35}
	employees["lucy"] = &Employee{"E100002", "Lucy", 40}
	fmt.Println(employees)
}

// Index func.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "🎉 Welcome!\n")
}

// GetEmployeeByName func.
func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var (
		ok   bool
		emp  *Employee
		info []byte
		err  error
	)
	if emp, ok = employees[name]; !ok {
		fmt.Println(emp, ok)
		w.Write([]byte("{\"error\":\"Not found\"}"))
		return
	}
	if info, err = json.Marshal(emp); err != nil {
		fmt.Println(info, err)
		w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err)))
		return
	}
	fmt.Println(info)
	w.Write(info)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/employee/:name", GetEmployeeByName)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
