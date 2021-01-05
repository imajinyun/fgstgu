package common

import "net/rpc"

// HelloWorldServiceName definition.
const HelloWorldServiceName = "server.HelloWorldService"

// HelloWorldServiceInterface interface.
type HelloWorldServiceInterface = interface {
	HelloWorld(request string, reponse *string) error
}

// RegisterHelloWorldService return a registered service.
func RegisterHelloWorldService(service HelloWorldServiceInterface) error {
	return rpc.RegisterName(HelloWorldServiceName, service)
}
