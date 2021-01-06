package main

import (
	"context"
	"fgstgu/go/src/b/grpc/b12/hello"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// HelloServiceImpl struct.
type HelloServiceImpl struct{}

// Hello method for HelloServiceImpl.
func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.String) (*hello.String, error) {
	reply := &hello.String{Value: "🎉 Server Response Hello: " + args.GetValue()}
	return reply, nil
}

// Channel method for HelloServiceImpl.
func (p *HelloServiceImpl) Channel(stream hello.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &hello.String{Value: "🎉 Server Response Hello: " + args.GetValue()}
		if err = stream.Send(reply); err != nil {
			return err
		}
	}
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
