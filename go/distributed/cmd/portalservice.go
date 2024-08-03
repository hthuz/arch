package cmd

import (
	"context"
	"distributed/portal"
	"distributed/registry"
	"distributed/service"
	"log"
)

func PortalServiceMain() {
	r := registry.Registration{
		ServiceName:      "Portal Service",
		ServiceHost:      "localhost",
		ServicePort:      "6000",
		ServiceURL:       "http://localhost:6000",
		RequiredServices: []string{"Log Service", "Grade Service"},
		ServiceUpdateURL: "http://localhost:6000/services",
		HeartbeatURL:     "http://localhost:6000/heartbeat",
	}

	ctx, err := service.Start(context.Background(), r, portal.RegisterHandlers)
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
	service.Stop(r)
}
