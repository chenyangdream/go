package main

import "fmt"
import "github.com/chenyangdream/stringutil"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf(stringutil.Reverse("Hello, world.\n"))
}
