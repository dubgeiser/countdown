package server

import (
	"countdown/internal/shared"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	shared.RenderTemplate(w, "layout.html", nil)
}
