package services

import (
	"app/pb"
	"context"
	"fmt"
	"time"
)

// obtido de pb/user_grpc.pb.go
// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
//  AddUserDetails(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserDetailsClient, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	// serve para caso adicionarmos algum servico inexistente no protocol buffers, o servidor não ira dar problema
	pb.UnimplementedUserServiceServer
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
