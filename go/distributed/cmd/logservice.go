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
		ServiceName: "Log Service",
		ServiceHost: "localhost",
		ServicePort: "3000",
	}
	ctx, err := service.Start(context.Background(), r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
}
