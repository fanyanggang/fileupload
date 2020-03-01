package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func main() {

	info, err := ioutil.ReadFile("./taskid.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	taskid := strings.Replace(string(info), "\n", "", -1)

	filepath := "./coverage_auto.out"
	bodyBuf := &bytes.Buffer{}
	fmt.Print("-----------------fileupload begin-----------------\n")

	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", filepath)
	if err != nil {
		fmt.Println("error writing to buffer:%v", err)
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Open panci:%v", err)
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Printf("copy err:%v\n", err)

		return
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.WriteField("taskid", string(taskid))
	bodyWriter.WriteField("type", "ec")
	bodyWriter.WriteField("apiName", "upload_ec")
	bodyWriter.Close()
	fmt.Printf("taskid:%v\n", string(taskid))

	res, err := http.Post("http://cloudwebqa.api.qq.com:8080", contentType, bodyBuf)
	if err != nil {
		fmt.Printf("http panci:%v, :%v", file, err)
		panic(err)
	}
	defer res.Body.Close()

	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("succ :%v\n", string(message))
	fmt.Print("-----------------fileupload end-----------------\n")

}
