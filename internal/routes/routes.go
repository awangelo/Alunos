package routes

import (
	"alunos/internal/handlers"
	"alunos/internal/services"
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

	mux.Handle("GET /alunos", services.AuthMiddleware(http.HandlerFunc(handlers.Alunos)))
	mux.Handle("GET /alunos/inserir", services.AuthMiddleware(http.HandlerFunc(handlers.InserirAlunoForm)))
	mux.Handle("POST /alunos/inserir", services.AuthMiddleware(http.HandlerFunc(handlers.InserirAluno)))
	mux.Handle("DELETE /alunos/{ra}", services.AuthMiddleware(http.HandlerFunc(handlers.RemoverAluno)))

	return mux
}
