package main

import (
	"fgstgu/go/src/b/grpc/b14/gen"
	"fgstgu/go/src/b/grpc/b14/hello"
	"log"
	"net/http"
)

// HelloService struct.
type HelloService struct{}

// Hello method for HelloService.
func (s *HelloService) Hello(request *gen.String, response *gen.String) error {
	response.Value = "🎉 Hello: " + request.GetValue()
	return nil
}

func main() {
	router := hello.HelloServiceHandler(new(HelloService))
	log.Fatal(http.ListenAndServe(":8080", router))
}
