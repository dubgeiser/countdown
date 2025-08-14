package numbers

import (
	"countdown/internal/shared"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	target, selection := Pick(2)
	shared.RenderTemplate(w, "layout.html", TplVarsNumbers{
		Game:      "Numbers",
		Target:    target,
		Selection: selection,
	})
}

