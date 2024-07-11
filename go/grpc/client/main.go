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

	req := &pb.SearchRequest{Name: "John"}
	// Search is the rpc defined in proto service, which receives SearchRequest and returns SearchResponse
	resp, _ := client.Search(context.Background(), req)

	fmt.Println(resp)

}
