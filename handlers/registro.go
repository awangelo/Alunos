package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Registro",
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/registro.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}
