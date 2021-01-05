package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// -> nc -l 1234
// -> go run client/main.go
// Output: {"method":"HelloWorldService.HelloWorld","params":["Client HelloWorldService request❗️"],"id":0}
//
// -> nc -l 1234
// -> echo -e '{"method":"HelloWorldService.Hello", "params":["HelloWorld"], "id":1001}' | nc localhost 1234
// Output: {"method":"HelloWorldService.Hello", "params":["HelloWorld"], "id":1001}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var (
		response string
		method   = "HelloWorldService.HelloWorld"
		request  = "Client HelloWorldService request❗️"
	)
	if err = client.Call(method, request, &response); err != nil {
		log.Fatal("Call error:", err)
	}
	fmt.Println(response)
}
