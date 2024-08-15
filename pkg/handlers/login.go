package handlers

import (
	"alunos/pkg/services"
	"html/template"
	"net/http"
)

var loginTemplate = template.Must(template.ParseFiles("layout.html", "login.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	// Parse dos templates.
	// Executa o template sem erro.
	loginTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Login",
		"Error": nil,
	})
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Pega os valores do formulario.
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Se o login for invalido, renderiza envia uma mensagem de erro.
		if !services.ValidateLogin(username, password) {
			loginTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{
				"Title": "Login",
				"Error": "Usuário ou senha incorretos.",
			})
			return
		}

		http.Redirect(w, r, "/alunos", http.StatusFound)
		return
	}
}
