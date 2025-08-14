package shared

import (
	"net/http"
	"text/template"
)

type Config struct {
	TemplatesGlob string
}

var config Config = Config{
	TemplatesGlob: "templates/*.html",
}
var allTemplates = template.Must(template.ParseGlob(config.TemplatesGlob))

func RenderTemplate(w http.ResponseWriter, tpl string, data any) {
	err := allTemplates.ExecuteTemplate(w, tpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
