package models

import (
	"alunos/internal/services"
	"encoding/json"
	"net/http"
)

func UserHasCookie(w http.ResponseWriter, r *http.Request) bool {
	// Verifica se o usuário já possui um token de sessão válido.
	oldSessionToken, err := r.Cookie("session_token")
	if err == nil && services.IsValidSession(oldSessionToken.Value) {
		// Retorna uma resposta JSON com status OK indicando sucesso.
		userIsAuthenticated(w)
		return true
	}

	return false
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
