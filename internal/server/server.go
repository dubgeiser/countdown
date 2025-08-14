package server

import (
	"countdown/internal/letters"
	"countdown/internal/numbers"
	"countdown/internal/shared"
	"fmt"
	"log"
	"net/http"
)

func StartServer(port int) {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/numbers/", numbers.Index)
	http.HandleFunc("/letters/", letters.Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	shared.RenderTemplate(w, "layout.html", nil)
}
