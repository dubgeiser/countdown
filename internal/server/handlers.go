package server

import (
	"countdown/internal/view"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "", view.ViewData{
		ShowGameNavigation: true,
	})
}
