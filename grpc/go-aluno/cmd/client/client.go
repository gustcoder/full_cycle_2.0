package main

import (
	"app/pb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//AddUserDetails(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
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

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Aluno 1",
			Email: "aluno1@fullcycle.com",
		},
		{
			Id:    "2",
			Name:  "Aluno 2",
			Email: "aluno2@fullcycle.com",
		},
		{
			Id:    "3",
			Name:  "Aluno 3",
			Email: "aluno3@fullcycle.com",
		},
		{
			Id:    "4",
			Name:  "Aluno 4",
			Email: "aluno4@fullcycle.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving user: %v", err)
				break
			}
			fmt.Printf("Receiving user %v with status %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Aluno 1",
			Email: "aluno1@fullcycle.com",
		},
		{
			Id:    "2",
			Name:  "Aluno 2",
			Email: "aluno2@fullcycle.com",
		},
		{
			Id:    "3",
			Name:  "Aluno 3",
			Email: "aluno3@fullcycle.com",
		},
		{
			Id:    "4",
			Name:  "Aluno 4",
			Email: "aluno4@fullcycle.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error sending Users stream:  %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error on response from Users stream:  %v", err)
	}

	fmt.Println(res)
}
