package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go performTask(ctx)

	select {
	case <-ctx.Done():
		fmt.Printf("%T\n", ctx)
		fmt.Println(ctx)
		fmt.Printf("%T\n", ctx.Done())
		fmt.Println(ctx.Done())
		fmt.Println("Task timeout")
	}
}

// If performTask is not finished within 2 seconds, the program terminates
func performTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed")
	}

}
