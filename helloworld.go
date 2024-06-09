package main

import (
	"context"
	"fmt"
	"github.com/GRPCPractice/proto/proto/helloworld"
	"io"
	"time"
)

func SayHello(c helloworld.GreeterClient) {
	fmt.Println("@@@ Start SayHello")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: defaultName})
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}
	fmt.Printf("Greeting: %s\n", r.GetMessage())
	fmt.Println("@@@ End SayHello")
}

func StreamHelloRequests(c helloworld.GreeterClient) {
	fmt.Println("@@@ Start StreamHelloRequests")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.StreamHelloRequests(ctx)
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}

	for i := 0; i < 10; i++ {
		err := r.Send(&helloworld.HelloRequest{Name: defaultName + fmt.Sprint(i)})
		if err != nil {
			fmt.Printf("could not greet: %v\n", err)
			break
		}
	}

	reply, err := r.CloseAndRecv()
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}

	fmt.Printf("Greeting: %s\n", reply.GetMessage())
	fmt.Println("@@@ End StreamHelloRequests")
}

func StreamHelloReplies(c helloworld.GreeterClient) {
	fmt.Println("@@@ Start StreamHelloReplies")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.StreamHelloReplies(ctx, &helloworld.HelloRequest{Name: defaultName})
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}

	for {
		reply, err := r.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("could not greet: %v\n", err)
			break
		}

		fmt.Printf("Greeting: %s\n", reply.GetMessage())
	}
	fmt.Println("@@@ End StreamHelloReplies")
}

func SayHelloChat(c helloworld.GreeterClient) {
	fmt.Println("@@@ Start SayHelloChat")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHelloChat(ctx)
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}

	for i := 0; i < 10; i++ {
		err := r.Send(&helloworld.HelloRequest{Name: defaultName + fmt.Sprint(i)})
		if err != nil {
			fmt.Printf("could not greet: %v\n", err)
			break
		}

		reply, err := r.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("could not greet: %v\n", err)
			break
		}

		fmt.Printf("Greeting: %s\n", reply.GetMessage())
	}

	if err := r.CloseSend(); err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}
	fmt.Println("@@@ End SayHelloChat")
}
