package server

import (
	"countdown/internal/letters"
	"countdown/internal/numbers"
	"fmt"
	"log"
	"net/http"
)

func StartServer(port int) {
	http.HandleFunc("/", Index)
	http.HandleFunc("/numbers/", numbers.Index)
	http.HandleFunc("/letters/", letters.Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
