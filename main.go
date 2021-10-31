package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var port = flag.Int("port", 8888, "port")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = append(args, ".")
	}
	log.Println("args: " + args[0])
	folderPath := args[0]
	randomPortIfInUse(port)

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

	log.Println(fmt.Sprintf("listen on port %d, folder:%s", *port, folderPath))
	log.Println(fmt.Sprintf("http://localhost:%d", *port))

	if err := http.ListenAndServe(":"+strconv.Itoa(*port), nil); err != nil {
		log.Fatal(err)
	}
}

func randomPortIfInUse(port *int) {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		fmt.Println("Port is in use: " + strconv.Itoa(*port))
		*port = generatePortBetween(3000, 4000)
		fmt.Println("Switch to port: " + strconv.Itoa(*port))

		randomPortIfInUse(port)
	} else {
		ln.Close()
	}
}

func generatePortBetween(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

type FilePart struct {
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}
