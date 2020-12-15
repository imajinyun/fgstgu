package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baidu = "https://baidu.com/"
	tmall = "https://tmall.com/"
	meitu = "https://meitu.com/"
)

func main() {
	go responseSize(baidu)
	go responseSize(tmall)
	go responseSize(meitu)
}

func responseSize(url string) int {
	fmt.Println("Getting:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return len(body)
}
