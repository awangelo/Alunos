package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Alunos(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Alunos",
	}

	tmpl, err := template.ParseFiles("web/templates/layout.html", "web/templates/alunos.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}
