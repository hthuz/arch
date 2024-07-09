package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Person struct {
	Name    string
	Gender  string
	Age     int
	Friends []*Person
}
type GetPersonArgs struct {
	Name string
}

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dailing: ", err)
	}
	args := GetPersonArgs{Name: "Mike"}
	var reply Person
	err = client.Call("Server.GetInfo", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

}
