package main

import "fmt"
import "path"

func main() {
	fmt.Println(path.Clean("././././hello/./hhlo.ssl"))
}
