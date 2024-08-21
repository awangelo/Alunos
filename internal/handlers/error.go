package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/error.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
