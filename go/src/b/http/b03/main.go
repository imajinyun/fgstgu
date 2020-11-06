package main

import (
	"fmt"
	"net/http"
)

func main() {
	sendGetRequest()
}

func sendGetRequest() {
	url := "https://www.toutiao.com/search/suggest/initial_page/"
	client := http.Client{}
	response, err := client.Get(url)
	checkError(err)
	if response.StatusCode == 200 {
		fmt.Printf("✔️ Resquest successfully!")
		defer response.Body.Close()
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
