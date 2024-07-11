package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"net"

	"google.golang.org/grpc"
)

// Your implemented server
type server struct {
	pb.UnimplementedSearchServiceServer
}

// Should be the same name as defined in proto service
func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	name := in.Name
	persons := []pb.Person{
		{Name: "Mike", Age: 15},
		{Name: "John", Age: 13},
		{Name: "Sarah", Age: 14},
	}

	fmt.Println("Get Req", *in)
	var resp pb.SearchResponse
	for _, person := range persons {
		if person.Name == name {
			resp.Person = &person
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
