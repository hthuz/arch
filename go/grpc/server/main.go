package main

import (
	"context"
	"errors"
	"fmt"
	"grpc/pb"
	"net"

	"google.golang.org/grpc"
)

// Your implemented server
type server struct {
	pb.UnimplementedSearchServiceServer
}

var persons = []pb.Person{
	{Id: 1, Name: "Mike", Age: 15},
	{Id: 2, Name: "John", Age: 13},
	{Id: 3, Name: "Sarah", Age: 14},
	{Id: 4, Name: "Joe", Age: 15},
}

// Should be the same name as defined in proto service
func (s *server) SearchName(ctx context.Context, in *pb.SearchNameRequest) (*pb.SearchNameResponse, error) {
	name := in.Name
	fmt.Println("Get Req", *in)
	var resp pb.SearchNameResponse
	resp.Persons = []*pb.Person{}
	for _, person := range persons {
		if person.Name == name {
			resp.Persons = append(resp.Persons, &person)
		}
	}
	return &resp, nil
}

func (s *server) SearchAge(ctx context.Context, in *pb.SearchAgeRequest) (*pb.SearchAgeResponse, error) {
	age := in.Age
	// This error will be returned as rpc error
	if age <= 0 {
		return nil, errors.New("Age can't be smaller than 0")
	}
	fmt.Println("Get Req", *in)
	var resp pb.SearchAgeResponse
	resp.Persons = []*pb.Person{}
	for _, person := range persons {
		if person.Age == age {
			resp.Persons = append(resp.Persons, &person)
		}
	}
	return &resp, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":8000")

	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, &server{})

	s.Serve(listen)

}
