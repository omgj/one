package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func yi(w http.ResponseWriter, r *http.Request) {
	log.Println("here")
	b, _ := ioutil.ReadFile("index.html")
	io.WriteString(w, string(b))
}

func main() {
	http.Handle("/stuff/", http.StripPrefix("/stuff/", http.FileServer(http.Dir("stuff"))))
	http.HandleFunc("/", yi)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
