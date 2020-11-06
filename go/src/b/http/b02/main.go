package main

import (
	"fmt"
	"net/http"
)

func main() {
	sendHttpRequest()
}

func sendHttpRequest() {
	url := "https://gate.lagou.com/v1/entry/privacyPolicy/query"
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	check(err)
	cookie := &http.Cookie{Name: "username", Value: "Peter"}
	request.AddCookie(cookie)
	response, err := client.Do(request)
	request.Header.Set("Accept-Language", "zh-CN")
	defer response.Body.Close()
	fmt.Printf("Header: %+v\n", request.Header)
	fmt.Printf("Response status code: %v\n", response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println("Request successfully!")
		check(err)
	} else {
		fmt.Println("Request failed!")
	}
}

func check(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("An exception has occurred: ", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
