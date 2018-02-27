package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:4002/ob/profile"

	fmt.Println("Get url = ", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP Get ", url, "error")
		return
	}

	io.Copy(os.Stdout, res.Body)
}
