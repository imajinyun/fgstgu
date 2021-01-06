package main

import (
	"context"
	"fgstgu/go/src/b/grpc/b13/api"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := api.NewPublishServiceClient(conn)
	if _, err := client.Publish(
		context.Background(),
		&api.String{Value: "golang: hello golang"}); err != nil {
		log.Fatal(err)
	}
	if _, err = client.Publish(
		context.Background(),
		&api.String{Value: "docker: hello docker"}); err != nil {
		log.Fatal(err)
	}
}
