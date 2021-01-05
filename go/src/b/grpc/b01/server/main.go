package main

import (
	"log"
	"net"
	"net/rpc"
)

// HelloWorldService struct.
type HelloWorldService struct{}

// HelloWorld returns a string.
func (hw *HelloWorldService) HelloWorld(request string, response *string) error {
	*response = "🎉 Hello: " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloWorldService", new(HelloWorldService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
