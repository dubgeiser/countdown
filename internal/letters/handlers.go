package letters

import (
	"countdown/internal/view"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "letters.html", view.ViewData{
		Game: "Letters",
	})
}

func PickVowel(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "letters.html", view.ViewData{
		Game:         "Letters",
		PickedLetter: 'O',
	})
}

func PickConsonant(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "letters.html", view.ViewData{
		Game:         "Letters",
		PickedLetter: 'L',
	})
}
