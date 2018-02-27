package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for {
		i := <-ch
		fmt.Println("Value received: ", i)
	}
}
