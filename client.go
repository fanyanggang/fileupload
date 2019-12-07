package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"mime/multipart"
)

func main() {

	total := "/Users/fengqingyang/tempfile"
	files, err := ioutil.ReadDir(total)
	if err != nil{
		fmt.Errorf("read file err:%v", err)
	}
	fmt.Printf("beigin:%v\n", files)
	for _, file := range files{
		fileName := file.Name()
		filepath.Walk(fileName, func(path string, info os.FileInfo, err error)error{

            absolutePath :=total + "/" + fileName
			fmt.Printf("absolutePath:%v\n", absolutePath)
			
			bodyBuf := &bytes.Buffer{}
			fmt.Printf("filePahtq:%v\n", bodyBuf)

			bodyWriter := multipart.NewWriter(bodyBuf)
			fileWriter, err := bodyWriter.CreateFormFile("file", filepath.Base(absolutePath))
			if err != nil {
				fmt.Println("error writing to buffer:%v", err)
				return err
			}
	
			file, err := os.Open(absolutePath)
			if err != nil {
				fmt.Printf("Open panci:%v", err)
				panic(err)
			}
			defer file.Close()

			_, err = io.Copy(fileWriter, file)
			if err != nil {
				return err
			}
			contentType := bodyWriter.FormDataContentType()
			bodyWriter.Close()

			res, err := http.Post("http://127.0.0.1:5050/upload", contentType, bodyBuf)
			if err != nil {
				fmt.Printf("http panci:%v,   :%v", file, err)
				panic(err)
			}
			defer res.Body.Close()

			message, _ := ioutil.ReadAll(res.Body)
			fmt.Printf("succ :%v\n", string(message))
			return nil
			})
	}
}
