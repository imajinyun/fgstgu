package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	. "fgstgu/go/src/b/grpc/b02/common"
)

// HelloWorldService struct.
type HelloWorldService struct{}

// HelloWorld return a string.
func (hw *HelloWorldService) HelloWorld(request string, response *string) error {
	*response = "🎉 Server HelloWorldService response: " + request
	fmt.Println(fmt.Sprintf("%s", *response))
	return nil
}

func main() {
	RegisterHelloWorldService(new(HelloWorldService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
