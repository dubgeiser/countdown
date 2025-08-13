package server

import (
	"countdown/internal/numbers"
	"fmt"
	"html/template"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "layout.html", nil)
}

func NumbersHandler(w http.ResponseWriter, r *http.Request) {
	target, selection := numbers.Pick(2)
	renderTemplate(w, "layout.html", NumbersData{
		Game:      "Numbers",
		Target:    target,
		Selection: selection,
	})
}

func LettersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Letters: %s", r.URL.Path[1:])
}

