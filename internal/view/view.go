package view

import (
	"html/template"
	"net/http"
	"strings"
)

type ViewData struct {
	Game               string
	Target             []int
	Selection          []int
	BigNumberPicker    []int
	ShowGameNavigation bool
	LettersSelection   []rune
	PickedLetter       rune
}

var tplBase = template.Must(template.ParseFiles("templates/layout.html"))

func check(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	return true
}

// Leave `name` empty if no extra template must be rendered.
func Render(w http.ResponseWriter, tplFile string, data ViewData) {
	var tpl *template.Template
	var err error
	if len(tplFile) > 0 {
		tpl, err = template.Must(tplBase.Clone()).ParseFiles(
			strings.Join([]string{"templates", tplFile}, "/"))
	} else {
		tpl = template.Must(tplBase.Clone())
	}
	if !check(err, w) {
		return
	}
	check(tpl.ExecuteTemplate(w, "layout.html", data), w)
}
