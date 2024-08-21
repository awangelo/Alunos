package routes

import (
	"alunos/internal/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("GET /login", handlers.Login)
	mux.HandleFunc("POST /login", handlers.LoginAuth)
	mux.HandleFunc("GET /error", handlers.Error)

	mux.Handle("GET /alunos", handlers.AuthMiddleware(http.HandlerFunc(handlers.Alunos)))
	mux.Handle("GET /alunos/inserir", handlers.AuthMiddleware(http.HandlerFunc(handlers.InserirAlunoForm)))
	mux.Handle("POST /alunos/inserir", handlers.AuthMiddleware(http.HandlerFunc(handlers.InserirAluno)))
	mux.Handle("DELETE /alunos/{ra}", handlers.AuthMiddleware(http.HandlerFunc(handlers.RemoverAluno)))

	return mux
}
