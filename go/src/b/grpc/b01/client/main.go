package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}
	var response string
	if err := client.Call("HelloWorldService.HelloWorld", "😂", &response); err != nil {
		log.Fatal("Call error:", err)
	}
	fmt.Println(response)
}
