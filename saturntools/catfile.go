package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fileHash := "zb2rhi36Gc9GJWijLEL6zW45MBux5FcFv5gJmjXA7VAMozEXY"
	if len(os.Args) > 1 {
		fileHash = os.Args[1]
	}
	url := "http://localhost:4002/ob/cat" + "/" + fileHash

	fmt.Println("Get url = ", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP Get ", url, "error")
		return
	}

	fmt.Println("Server return file content:")
	io.Copy(os.Stdout, res.Body)
}
