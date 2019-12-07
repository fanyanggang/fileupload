package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	data, fileHeader, err := r.FormFile("file")
	fmt.Printf("data:%v\n", data)
	fmt.Printf("fileHeader:%v\n", fileHeader.Header)
	file, err := os.Create("./" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(file, data)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("upload success"))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":5050", nil)
}