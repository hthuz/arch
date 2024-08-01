package cmd

import (
	"context"
	"distributed/registry"
	"distributed/service"
	"log"
)

func RegistryServiceMain() {
	// Registry Service will register itself first
	r := registry.Registration{
		ServiceName: "Registry Service",
		ServiceHost: "localhost",
		ServicePort: "5000",
	}
	ctx, err := service.Start(context.Background(), r, registry.RegisterHandlers)
	if err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
}
