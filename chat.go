package main

import (
	"context"
	"fmt"
	"github.com/GRPCPractice/proto/proto/chat"
	"time"
)

func Connect(c chat.ChatServiceClient) {
	fmt.Println("@@@ Start Connect")
	ctx := context.Background()
	stream, err := c.Connect(ctx, &chat.ConnectRequest{UserId: "1"})
	if err != nil {
		fmt.Printf("could not connect: %v\n", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				ctx.Done()
				fmt.Printf("could not receive: %v\n", err)
				break
			}
			fmt.Printf("Receive: %s\n", msg.GetMessage())
		}
	}()

	fmt.Println("@@@ End Connect")
}

func Send(c chat.ChatServiceClient) {
	fmt.Println("@@@ Start Send (Type 'exit' to exit)")
	for {
		var message string
		fmt.Scan(&message)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		if message == "exit" {
			break
		}

		msg := &chat.ChatMessage{
			UserId:  "1",
			Message: message,
		}

		_, err := c.Send(ctx, msg)
		if err != nil {
			fmt.Printf("could not send: %v\n", err)
			break
		}
		cancel()
	}

	fmt.Println("@@@ End Send")
}
