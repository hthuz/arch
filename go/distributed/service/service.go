package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
	"time"
)

// register service and call ListenAndService
func Start(ctx context.Context, r registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, r)

	// HOLDER, wait for server to start
	if r.ServiceName == "Registry Service" {
		time.Sleep(1 * time.Second)
	}

	// Register to registry service
	err := registry.RegisterService(r)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, r registry.Registration) context.Context {
	// Create a child context of input ctx
	ctx, cancel := context.WithCancel(ctx)

	// HTTP server
	var srv http.Server
	srv.Addr = ":" + r.ServicePort

	fmt.Printf("%s started, input <exit> to exit\n", r.ServiceName)
	// Stop by user interrupt
	go func() {
		var s string
		for s != "exit" {
			fmt.Scanln(&s)
		}
		// Unregister from registry service
		if err := registry.UnregisterService(r.GetURL()); err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	go func() {
		// Stop by server error
		err := srv.ListenAndServe()
		fmt.Println(err)
		// Unregister from registry service
		if err := registry.UnregisterService(r.GetURL()); err != nil {
			log.Println(err)
		}
		cancel()
	}()

	return ctx // Return context for propagation
}
