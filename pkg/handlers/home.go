package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	homeTemplate := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/home.html"))
	homeTemplate.Execute(w, map[string]interface{}{
		"Title": "Inicio",
	})
}
