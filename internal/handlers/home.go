package handlers

import (
	"alunos/internal/models"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	numAlunos := models.GetNumAlunos()

	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/home.html"))
	if err := tmpl.Execute(w, numAlunos); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
