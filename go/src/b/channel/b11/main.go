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
	tests = "https://tests.com"
)

func main() {
	sizeChannel := make(chan int)
	go responseSize(baidu, sizeChannel)
	go responseSize(tmall, sizeChannel)
	go responseSize(meitu, sizeChannel)
	go responseSize(tests, sizeChannel)
	fmt.Println(<-sizeChannel)
	fmt.Println(<-sizeChannel)
	fmt.Println(<-sizeChannel)
	fmt.Println(<-sizeChannel)
}

func responseSize(url string, channel chan int) {
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
	channel <- len(body)
}
