package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("sample/html"))

	http.Handle("/", fs)

	log.Println("listen on port 8888")
	log.Println("http://localhost:8888")

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
