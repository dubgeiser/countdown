package shared

import (
	"net/http"
	"text/template"
)

var allTemplates = template.Must(template.ParseGlob("templates/*.html"))

func RenderTemplate(w http.ResponseWriter, tpl string, data any) {
	err := allTemplates.ExecuteTemplate(w, tpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
