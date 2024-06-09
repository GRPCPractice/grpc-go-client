package main

import (
	"fmt"
	"github.com/GRPCPractice/proto/proto/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "BeomJun"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		panic(err)
	}

	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	SayHello(c)
	StreamHelloRequests(c)
	StreamHelloReplies(c)
	SayHelloChat(c)
}
