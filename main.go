package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	folderPath := flag.String("folder", ".", "folder path")
	flag.Parse()
	log.Println("folder: " + *folderPath)

	fs := http.FileServer(http.Dir(*folderPath))

	http.Handle("/", fs)

	log.Println("listen on port 8888, folderPath:" + *folderPath)
	log.Println("http://localhost:8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
