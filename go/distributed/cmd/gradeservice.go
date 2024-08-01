package cmd

import (
	"context"
	"distributed/grades"
	"distributed/registry"
	"distributed/service"
	"log"
)

func GradeServiceMain() {
	r := registry.Registration{
		ServiceName: "Grade Service",
		ServiceHost: "localhost",
		ServicePort: "4000",
	}
	ctx, err := service.Start(context.Background(), r, grades.RegisterHandlers)
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}
