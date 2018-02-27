package main

import "fmt"
import "time"

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Monitor quit, stop...")
				return
			default:
				fmt.Println("goroutine moitoring...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("notify monitor to stop")
	stop <- true
	time.Sleep(5 * time.Second)
}
