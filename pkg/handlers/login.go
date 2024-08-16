package handlers

import (
	"alunos/pkg/services"
	"encoding/json"
	"html/template"
	"net/http"
)

var loginTemplate = template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/login.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	// Executa o template sem erro.
	loginTemplate.Execute(w, nil)
}

// LoginAuth vai devolver json com erro se o login for invalido.
func LoginAuth(w http.ResponseWriter, r *http.Request) {
	// Pega os valores do formulario.
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Se o login for invalido, renderiza envia uma mensagem de erro.
	if !services.ValidateLogin(username, password) {
		// Escreve o header e o status code no writer.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Usuário ou senha incorretos.",
		})
		return
	}

	// Verifica se o usuário já possui um token de sessão válido.
	oldSessionToken, err := r.Cookie("session_token")
	if err == nil && services.IsValidSession(oldSessionToken.Value) {
		// Retorna uma resposta JSON com status OK indicando sucesso.
		userIsAuthenticated(w)
		return
	}

	// Token `string`.
	sessionToken := services.GenerateSessionToken()

	if !services.SaveSessionToken(username, sessionToken) {
		// Escreve o header e o status code no writer.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Erro ao salvar o token de sessão.",
		})
		return
	}

	// Adciona `Set-Cookie` ao header do writer.
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",  // Sera usado apenas nas rotas de alunos.
		HttpOnly: true, // Impede que o cookie seja acessado por js.
	})

	userIsAuthenticated(w)
}

// Retorna uma resposta JSON com status OK indicando sucesso.
func userIsAuthenticated(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Login bem-sucedido.",
		"redirect": "/alunos",
	})
}
