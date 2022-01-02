package main

import (
	"app/pb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error when Dial Connection: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	AddUserDetails(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Gus",
		Email: "gust@client.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Error AddUser: %v", err)
	}

	fmt.Println(res)
}

func AddUserDetails(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Gus",
		Email: "gust@client.com",
	}

	responseStream, err := client.AddUserDetails(context.Background(), req)
	if err != nil {
		log.Fatalf("Error AddUserDetails: %v", err)
	}

	for {
		stream, err := responseStream.Recv() // Receive
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error to receive Stream: %v", err)
		}

		fmt.Println(stream.Status)
	}
}
