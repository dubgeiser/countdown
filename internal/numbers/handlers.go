package numbers

import (
	"countdown/internal/shared"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	shared.RenderTemplate(w, "layout.html", ViewData{
		Game:      "Numbers",
		BigNumberPicker: []int{0,1,2,3,4},
	})
}

func Picker(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("big"))
	if err != nil || n < 0 || n > 4{
		http.Error(w, "Invalid amount of big numbers", 500)
		return
	}
	target, selection := Pick(n)
	shared.RenderTemplate(w, "layout.html", ViewData{
		Game:      "Numbers",
		Target:    target,
		Selection: selection,
	})
}
