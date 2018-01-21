package main

import (
	"fmt"
	"time"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}

	ch := make(chan bool, 1)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1e9)
		//timeout <- true
	}()

	select {
	case <-ch:
	case <-timeout:
	}
}
