package main

import (
	"log"
	"net/rpc"

	. "fgstgu/go/src/b/grpc/b02/common"
)

// HelloWorldServiceClient struct.
type HelloWorldServiceClient struct {
	*rpc.Client
}

// HelleWorldServiceInterface interface.
var _ HelloWorldServiceInterface = (*HelloWorldServiceClient)(nil)

// DialHelloWorldService return a HelloWorldServiceClient.
func DialHelloWorldService(network, address string) (*HelloWorldServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloWorldServiceClient{Client: client}, nil
}

// HelloWorld return a client call.
func (hw *HelloWorldServiceClient) HelloWorld(request string, response *string) error {
	return hw.Client.Call(HelloWorldServiceName+".HelloWorld", request, response)
}

func main() {
	client, err := DialHelloWorldService("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}
	var response string
	if err = client.HelloWorld("Client HelloWorldService request❗️", &response); err != nil {
		log.Fatal("Client error:", err)
	}
}
