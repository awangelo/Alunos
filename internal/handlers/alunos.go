package handlers

import (
	"alunos/internal/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Alunos(w http.ResponseWriter, r *http.Request) {
	alunos := models.GetAlunos()

	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/alunos.html"))
	err := tmpl.Execute(w, alunos)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func InserirAlunoForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/inserir.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func InserirAluno(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	ra := r.FormValue("ra")
	m1, m2, m3 := r.FormValue("m1"), r.FormValue("m2"), r.FormValue("m3")

	if email == "" || ra == "" || m1 == "" || m2 == "" || m3 == "" {
		http.Error(w, "Todos os campos são obrigatórios.", http.StatusBadRequest)
		return
	}

	err := models.InsertAluno(ra, email, m1, m2, m3)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/alunos", http.StatusSeeOther)
}

func RemoverAluno(w http.ResponseWriter, r *http.Request) {
	// Como o metodo vai DELETE, nao utilizara parametros na URL
	// entao `ra := r.URL.Query().Get("ra")` nao funcionara.
	parts := strings.Split(r.URL.Path, "/")
	ra := parts[len(parts)-1]

	err := models.DeleteAluno(ra)
	if err != nil {
		log.Fatal(err)
	}

	// Recarrega a pagina de alunos.
	http.Redirect(w, r, "/alunos", http.StatusSeeOther)
}
