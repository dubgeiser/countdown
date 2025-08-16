package numbers

import (
	"countdown/internal/view"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view.Render(w, "numbers.html", view.ViewData{
		Game:            "Numbers",
		BigNumberPicker: []int{0, 1, 2, 3, 4},
	})
}

func Picker(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("big"))
	if err != nil || n < 0 || n > 4 {
		http.Error(w, "Invalid amount of big numbers", 500)
		return
	}
	target, selection := Pick(n)
	view.Render(w, "numbers.html", view.ViewData{
		Game:      "Numbers",
		Target:    target,
		Selection: selection,
	})
}
