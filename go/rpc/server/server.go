package main

import (
	"fmt"
	"log"
	"net"
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

type Server struct {
	persons []*Person
}

func (s *Server) GetInfo(req GetPersonArgs, resp *Person) error {
	for _, person := range s.persons {
		if person.Name == req.Name {
			*resp = *person
			fmt.Println("returning: ", resp)
			return nil
		}
	}
	resp = nil
	return nil
}

func NewSvr(persons []*Person) *Server {
	server := &Server{persons: persons}
	return server
}

func main() {
	persons := []*Person{
		{Name: "Mike", Gender: "Male", Age: 11},
		{Name: "Newton"},
		{Name: "Amy", Gender: "Female"},
	}
	server := NewSvr(persons)

	rpc.Register(server)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)

}
