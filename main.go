package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, ".")
	}
	log.Println("args: " + args[0])
	folderPath := args[0]

	fs := http.FileServer(http.Dir(folderPath))

	http.Handle("/", fs)

	http.HandleFunc("/upload", func(rw http.ResponseWriter, r *http.Request) {
		filePartBody, _ := ioutil.ReadAll(r.Body)
		var filePart FilePart
		json.Unmarshal(filePartBody, &filePart)
		println("=================================fileName: " + filePart.FileName)
		println("===" + filePart.Content)

		filePart.Content = strings.Replace(filePart.Content, "data:application/octet-stream;base64,", "", 1)
		content, _ := base64.URLEncoding.DecodeString(filePart.Content)

		println("content: " + string(content))

		file, _ := os.OpenFile(filePart.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		file.Write(content)
	})

	log.Println("listen on port 8888, folder:" + folderPath)
	log.Println("http://localhost:8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}

type FilePart struct {
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}
