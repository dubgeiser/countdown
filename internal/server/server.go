package server

import (
	"countdown/internal/letters"
	"countdown/internal/numbers"
	"fmt"
	"log"
	"net/http"
)

func StartServer(port int) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/", Index)
	http.HandleFunc("/numbers/", numbers.Index)
	http.HandleFunc("/numbers/pick", numbers.Picker)
	http.HandleFunc("/letters/", letters.Index)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
