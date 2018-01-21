package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}

var arr2D = [][]int{{1, 2}, {3, 4}}

func main() {
	fmt.Println(arr)
	fmt.Println(arr2D)
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1.len = ", len(s1), "s1.cap = ", cap(s1))
	fmt.Println(s1[:1])
	fmt.Println(s1[3:])

	s2 := append(s1[:1], s1[2:]...)
	fmt.Println("s1.len = ", len(s1), "s1.cap = ", cap(s1))
	fmt.Println("s2.len = ", len(s2), "s2.cap = ", cap(s2))

	fmt.Println(s1)
	fmt.Println(s2)
}
