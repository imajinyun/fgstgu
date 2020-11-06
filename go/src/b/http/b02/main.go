package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	sendGetRequest()
}

func sendGetRequest() {
	url := "https://gate.lagou.com/v1/entry/privacyPolicy/query"
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	checkError(err)
	cookie := &http.Cookie{Name: "username", Value: "Peter"}
	request.AddCookie(cookie)
	response, err := client.Do(request)
	request.Header.Set("Accept-Language", "zh-CN")
	defer response.Body.Close()
	fmt.Printf("Header: %+v\n", request.Header)
	fmt.Printf("Response status code: %v\n", response.StatusCode)

	if response.StatusCode == 200 {
		data, err := ioutil.ReadAll(response.Body)
		fmt.Println("✔️ Request successfully!")
		checkError(err)
		fmt.Println("✔️ Results:")
		fmt.Println(string(data))
	} else {
		fmt.Println("✖️ Request failed!")
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
