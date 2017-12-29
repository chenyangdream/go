package main

import (
	"fmt"
	"github.com/chenyangdream/stringutil"
	"math"
)

var i, j int = 1, 2

func printPi() {
	fmt.Println(math.Pi)
}

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf(stringutil.Reverse("Hello, world.\n"))
	printPi()
	printPi()
	printPi()
}
