package cmd

import (
	"context"
	"distributed/grades"
	"distributed/registry"
	"distributed/service"
	stlog "log"
)

func GradeServiceMain() {
	r := registry.Registration{
		ServiceName:      "Grade Service",
		ServiceHost:      "localhost",
		ServicePort:      "4000",
		ServiceURL:       "http://localhost:4000",
		ServiceUpdateURL: "http://localhost:4000/services",
		HeartbeatURL:     "http://localhost:4000/heartbeat",
		// RequiredServices: []string{"Log Service"},
	}
	ctx, err := service.Start(context.Background(), r, grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	// Need required service
	// logProvider, err := registry.GetProvider("Log Service")
	// if err != nil {
	// 	stlog.Println("dependent log service not found")
	// 	stlog.Fatal(err)
	// }

	// log.SetClientLogger(logProvider, r.ServiceName)

	<-ctx.Done()
	service.Stop(r)

}
