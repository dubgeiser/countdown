package main

import (
	"log"
	"net/http"
	"countdown/internal/server"
)


func main() {
	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/numbers/", server.NumbersHandler)
	http.HandleFunc("/letters/", server.LettersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
