package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	addr := "https://www.toutiao.com/search/suggest/initial_page/"
	data := url.Values{
		"keyword": {"美国大选"},
	}
	sendPostRequest(addr, data.Encode())
}

func sendPostRequest(url, str string) {
	reader := strings.NewReader(str)
	response, err := http.Post(url, "application/x-www-form-urlencoded", reader)
	checkError(err)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		fmt.Println("✔️ Request successfully!")
		checkError(err)
	} else {
		fmt.Println("✖️ Request failed!")
		fmt.Println("✖️ Request status: ", response.Status)
	}
}

func checkError(err error) {
	defer func() {
		if res, ok := recover().(error); ok {
			fmt.Println("✖️ An exception has occurred: ", res.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
