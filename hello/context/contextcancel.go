package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "monitor 1")
	go watch(ctx, "monitor 2")
	go watch(ctx, "monitor 3")

	time.Sleep(10 * time.Second)
	fmt.Println("ok, notify monitor to quit")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "monitor quit, stop...")
			return
		default:
			fmt.Println(name, "goroutine monitoring...")
			time.Sleep(2 * time.Second)
		}
	}
}
