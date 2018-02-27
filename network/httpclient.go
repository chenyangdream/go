package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type FileInfo struct {
	FileName string `json:"filename"`
	FilePath string `json:"filepath"`
}

func main() {
	url := "http://localhost:4002/ob/add"

	var fileInfo = FileInfo{
		FileName: "helloworld.txt",
		FilePath: "/Users/chenyang/go/src/github.com/chenyangdream/network/helloworld.txt"}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(fileInfo)
	res, _ := http.Post(url, "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}
