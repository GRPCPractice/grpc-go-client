package main

import (
	"fmt"
	"github.com/GRPCPractice/proto/proto/chat"
	"github.com/GRPCPractice/proto/proto/helloworld"
	"github.com/GRPCPractice/proto/proto/user"
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

	u := user.NewUserServiceClient(conn)
	CreateUser(u)
	GetUser(u)
	UpdateUser(u)
	DeleteUser(u)
	ListUsers(u)

	ch := chat.NewChatServiceClient(conn)
	Connect(ch)
	Send(ch)
}
