package main

import (
	"fmt"
	"grpc/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)

	nameReq := &pb.SearchNameRequest{Name: "John"}
	// Search is the rpc defined in proto service, which receives SearchRequest and returns SearchResponse
	nameResp, _ := client.SearchName(context.Background(), nameReq)
	fmt.Println(nameResp)

	// Search age
	ageReq := &pb.SearchAgeRequest{Age: -1}
	ageResp, err := client.SearchAge(context.Background(), ageReq)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ageResp)

}
