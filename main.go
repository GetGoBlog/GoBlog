package main

import (
	"io"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to GoBlog!\n")
}

func main() {
	http.HandleFunc("/", MainPage)
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
