package services

import (
	"app/pb"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

// obtido de pb/user_grpc.pb.go
// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
//  AddUserDetails(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserDetailsClient, error)
//  AddUsers(UserService_AddUsersServer) error
//  AddUserStreamBoth(UserService_AddUserStreamBothServer) error
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	// serve para caso adicionarmos algum servico inexistente no protocol buffers, o servidor não ira dar problema
	pb.UnimplementedUserServiceServer
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})
		if err != nil {
			log.Fatalf("Error sendind stream to the client: %v", err)
		}
	}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserDetails(req *pb.User, stream pb.UserService_AddUserDetailsServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting user",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Process finished")

			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Error stream AddUsers: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("Adding Name: ", req.GetName())
	}
}
