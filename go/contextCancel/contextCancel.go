package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go task(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task terminated")
			return
		default:
			fmt.Println("Doing task")
			time.Sleep(300 * time.Millisecond)
		}
	}

}
