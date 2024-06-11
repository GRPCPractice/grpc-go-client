package main

import (
	"context"
	"fmt"
	"github.com/GRPCPractice/proto/proto/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetUser(c user.UserServiceClient) {
	fmt.Println("@@@ Start GetUser")
	ctx := context.Background()
	r, err := c.GetUser(ctx, &user.UserID{Id: "1"})
	if err != nil {
		fmt.Printf("could not get user: %v\n", err)
	}
	fmt.Printf("User: %v\n", r)
	fmt.Println("@@@ End GetUser")
}

func CreateUser(c user.UserServiceClient) {
	fmt.Println("@@@ Start CreateUser")
	ctx := context.Background()
	r, err := c.CreateUser(ctx, &user.CreateUserRequest{Name: "BeomJun", Email: "kaq6822@naver.com"})
	if err != nil {
		fmt.Printf("could not create user: %v\n", err)
	}

	fmt.Printf("User: %v\n", r)
	fmt.Println("@@@ End CreateUser")
}

func UpdateUser(c user.UserServiceClient) {
	fmt.Println("@@@ Start UpdateUser")
	ctx := context.Background()
	r, err := c.UpdateUser(ctx, &user.UpdateUserRequest{Id: "1", Name: "Jun", Email: "kaq6822@naver.com"})
	if err != nil {
		fmt.Printf("could not update user: %v\n", err)
	}

	fmt.Printf("User: %v\n", r)
	fmt.Println("@@@ End UpdateUser")
}

func DeleteUser(c user.UserServiceClient) {
	fmt.Println("@@@ Start DeleteUser")
	ctx := context.Background()
	r, err := c.DeleteUser(ctx, &user.UserID{Id: "1"})
	if err != nil {
		fmt.Printf("could not delete user: %v\n", err)
	}

	fmt.Printf("User: %v\n", r)
	fmt.Println("@@@ End DeleteUser")
}

func ListUsers(c user.UserServiceClient) {
	fmt.Println("@@@ Start ListUsers")
	ctx := context.Background()
	r, err := c.ListUsers(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Printf("could not list users: %v\n", err)
	}

	fmt.Printf("Users: %v\n", r)
	fmt.Println("@@@ End ListUsers")
}
