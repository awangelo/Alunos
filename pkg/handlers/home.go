package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Inicio",
	}

	tmpl, err := template.ParseFiles("web/templates/layout.html", "web/templates/home.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}
