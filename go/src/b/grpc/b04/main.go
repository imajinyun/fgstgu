package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// HelloWorldService struct.
type HelloWorldService struct{}

// HelloWorld return a string.
func (s *HelloWorldService) HelloWorld(request string, response *string) error {
	*response = "🎉 Server HelloWorldService response: " + request
	fmt.Println(fmt.Sprintf("%s", *response))
	return nil
}

// -> go run main.go
// -> curl localhost:1234/jsonrpc -X POST \
//    --data '{"method":"HelloWorldService.HelloWorld","params":["Hello World"],"id":1001}'
func main() {
	rpc.RegisterName("HelloWorldService", new(HelloWorldService))
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
