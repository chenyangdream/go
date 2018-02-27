package main

import (
	"net/http"
	"log"
	"io"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" " +
						" enctype\"multipart/form-data\">" +
			           	"Choose an image to upload: <intput name=\"image\" type=\"file\" />" +
			           	"<input type=\"submit\" value=\"Upload\" />" +
						"</from>")
		return
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
