package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "mointor1")
	go watch(valueCtx)

	time.Sleep(10 * time.Second)
	fmt.Println("ok, notify minitor to quit")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx value is ", ctx.Value(key), "monitor quit, stop...")
			return
		default:
			fmt.Println("ctx value is ", ctx.Value(key), "goroutine monitoring...")
			time.Sleep(2 * time.Second)
		}
	}
}
