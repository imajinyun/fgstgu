package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var urls = []string{
	"https://baidu.com/",
	"https://tmall.com/",
	"https://meitu.com/",
	"https://tests.com",
}

// Page struct.
type Page struct {
	url  string
	size int
}

func main() {
	sizeChannel := make(chan Page)
	for _, url := range urls {
		go responseSize(url, sizeChannel)
	}
	for i := 0; i < len(urls); i++ {
		page := <-sizeChannel
		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}

func responseSize(url string, channel chan Page) {
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
	channel <- Page{url: url, size: len(body)}
}
