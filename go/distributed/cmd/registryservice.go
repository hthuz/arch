package cmd

import (
	"context"
	"distributed/registry"
	"distributed/service"
	"log"
)

func RegistryServiceMain() {

	registry.SetupRegistryService()
	// Registry Service will register itself first
	r := registry.Registration{
		ServiceName:      "Registry Service",
		ServiceHost:      "localhost",
		ServicePort:      "5000",
		ServiceURL:       "http://localhost:5000",
		ServiceUpdateURL: "http://localhost:5000/services",
		HeartbeatURL:     "http://localhost:5000/heartbeat",
	}
	ctx, err := service.Start(context.Background(), r, registry.RegisterHandlers)
	if err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
	service.Stop(r)
}
