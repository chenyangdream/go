package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	pwd, _ := os.Getwd()
	filename := "helloworld.txt"
	filePath := pwd + "/" + filename

	var fileInfo = FileInfo{
		FileName: "helloworld.txt",
		FilePath: filePath,
	}

	fmt.Println("Add file ", fileInfo.FilePath)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(fileInfo)
	res, _ := http.Post(url, "application/json; charset=utf-8", b)
	fmt.Println("Server return file hash:")
	io.Copy(os.Stdout, res.Body)
	fmt.Println("")
}
