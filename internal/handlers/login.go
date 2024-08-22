package handlers

import (
	"alunos/internal/models"
	"alunos/internal/services"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/login.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// LoginAuth vai devolver json com erro se o login for invalido.
func LoginAuth(w http.ResponseWriter, r *http.Request) {
	// Pega os valores do formulario.
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		// Escreve o header e o status code no writer.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Usuário e senha são obrigatórios.",
		})
		return
	}

	// Se o login for invalido, renderiza envia uma mensagem de erro e retorna.
	if !services.ValidateLogin(username, password) {
		// Escreve o header e o status code no writer.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Usuário ou senha incorretos.",
		})
		return
	}

	// Autentica o user caso ele tenha o cookie.
	models.UserHasCookie(w, r)

	// Gera um novo cookie.
	sessionToken, err := services.GenerateSessionToken()
	if err != nil {
		// Escreve o header e o status code no writer.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Erro ao gerar o token de sessão.",
		})
		return
	}

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
