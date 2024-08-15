package handlers

import (
	"alunos/pkg/services"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// AuthMiddleware verifica se o usuário está autenticado
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica se o cookie de sessão esta presente
		cookie, err := r.Cookie("session_token")

		// Se nao estiver presente ou nao for valido, redireciona para a pagina de aviso.
		if err != nil || !services.IsValidSession(cookie.Value) {
			// Redireciona para a home se nao estiver autenticado.
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		// Chama o proximo handler se estiver autenticado.
		next.ServeHTTP(w, r)
	})
}
