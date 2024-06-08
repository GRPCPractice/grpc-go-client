package main

import (
	"context"
	"fmt"
	"github.com/GRPCPractice/proto/proto/helloworld"
	"google.golang.org/grpc"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "BeomJun"
)

func main() {
	fmt.Println("Hello, World!")

	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		panic(err)
	}

	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: defaultName})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}

	fmt.Printf("Greeting: %s", r.GetMessage())
}
