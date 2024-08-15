package handlers

import (
	"alunos/pkg/services"
	"html/template"
	"net/http"
)

var loginTemplate = template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/login.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	// Parse dos templates.
	// Executa o template sem erro.
	loginTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Login",
		"Error": nil,
	})
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
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

	// Token `string`.
	sessionToken := services.GenerateSessionToken()

	if !services.SaveSessionToken(username, sessionToken) {
		http.Error(w, "Erro ao criar sessao.", http.StatusInternalServerError)
		return
	}

	// Adciona `Set-Cookie` ao header do writer.
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/alunos", // Sera usado apenas nas rotas de alunos.
		HttpOnly: true,      // Impede que o cookie seja acessado por js.
	})

	// Redireciona para a pagina de alunos.
	http.Redirect(w, r, "/alunos", http.StatusFound)
}
