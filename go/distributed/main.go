package main

import (
	"distributed/cmd"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Format: exec cmd")
	}
	usercmd := os.Args[1]
	switch usercmd {
	case "log":
		cmd.LogServiceMain()
	case "grade":
		cmd.GradeServiceMain()
	case "registry":
		cmd.RegistryServiceMain()
	case "portal":
		cmd.PortalServiceMain()
	default:
		log.Fatal("invalid cmd")
	}
}
