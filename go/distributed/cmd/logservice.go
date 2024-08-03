package cmd

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	stlog "log"
)

func LogServiceMain() {
	log.Run("output.log")
	r := registry.Registration{
		ServiceName:      "Log Service",
		ServiceHost:      "localhost",
		ServicePort:      "3000",
		ServiceURL:       "http://localhost:3000",
		ServiceUpdateURL: "http://localhost:3000/services",
		HeartbeatURL:     "http://localhost:3000/heartbeat",
	}
	ctx, err := service.Start(context.Background(), r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	service.Stop(r)
}
