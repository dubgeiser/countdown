package main

import (
	"countdown/numbers"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Config struct {
	TemplatesGlob string
}

type NumbersData struct {
	Game      string
	Target    int
	Selection []int
}

var config Config = Config{
	TemplatesGlob: "templates/*.html",
}
var allTemplates = template.Must(template.ParseGlob(config.TemplatesGlob))

func renderTemplate(w http.ResponseWriter, tpl string, data any) {
	err := allTemplates.ExecuteTemplate(w, tpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "layout.html", nil)
}

func numbersHandler(w http.ResponseWriter, r *http.Request) {
	target, selection := numbers.Pick(2)
	renderTemplate(w, "layout.html", NumbersData{
		Game:      "Numbers",
		Target:    target,
		Selection: selection,
	})
}

func lettersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Letters: %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/numbers/", numbersHandler)
	http.HandleFunc("/letters/", lettersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
