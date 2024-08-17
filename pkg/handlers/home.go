package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}
