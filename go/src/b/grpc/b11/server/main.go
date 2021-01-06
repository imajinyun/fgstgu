package main

import (
	context "context"
	"fgstgu/go/src/b/grpc/b11/hello"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

// HelloServiceImpl struct.
type HelloServiceImpl struct{}

// Hello method for HelloServiceImpl.
func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.String) (*hello.String, error) {
	reply := &hello.String{Value: "🎉 Server Response Hello: " + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listen)
}
