package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// HelloWorldService struct.
type HelloWorldService struct{}

// HelloWorld return a string.
func (s *HelloWorldService) HelloWorld(request string, response *string) error {
	*response = "🎉 Server HelloWorldService response: " + request
	fmt.Println(fmt.Sprintf("%s", *response))
	return nil
}

// -> go run rpcserver/main.go
func main() {
	rpc.Register(new(HelloWorldService))
	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}
