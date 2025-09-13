package server

import (
	"countdown/internal/view"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "", viewData())
}

func viewData() view.ViewData {
	return view.ViewData{
		Menu: menu(),
	}
}

func menu() []view.Link {
	return []view.Link{
		{Href: "/numbers", Label: "Numbers"},
		{Href: "/letters", Label: "Letters"},
	}
}
