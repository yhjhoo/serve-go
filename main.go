package main

import (
	"flag"
	"log"
	"net/http"
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

	log.Println("listen on port 8888, folder:" + folderPath)
	log.Println("http://localhost:8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
