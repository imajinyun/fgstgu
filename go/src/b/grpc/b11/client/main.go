package main

import (
	"context"
	"fgstgu/go/src/b/grpc/b11/hello"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hello.NewHelloServiceClient(conn)
	reply, err := client.Hello(
		context.Background(),
		&hello.String{Value: "👌 Client Request Hello"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
