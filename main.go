package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/stuff/", http.StripPrefix("/stuff/", http.FileServer(http.Dir("stuff"))))
	http.HandleFunc("/", yi)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
