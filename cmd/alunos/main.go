package main

import (
	"alunos/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	// Novo mux para rotear as requisicoes, com os handlers definidos para cada metodo e rota.
	mux := http.NewServeMux()

	// Servir arquivos estaticos.
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("GET /login", handlers.Login)
	mux.HandleFunc("POST /login", handlers.LoginAuth)
	mux.HandleFunc("GET /error", handlers.Error)

	// Middleware de autenticacao para as rotas protegidas.
	mux.Handle("GET /alunos", handlers.AuthMiddleware(http.HandlerFunc(handlers.Alunos)))
	// mux.Handle("/alunos/inserir", handlers.AuthMiddleware(http.HandlerFunc(handlers.InserirAluno)))
	// mux.Handle("/alunos/remover", handlers.AuthMiddleware(http.HandlerFunc(handlers.RemoverAluno)))

	log.Printf("Listening on :80")
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("Error ao iniciar o servidor: %v", err)
	}
}
