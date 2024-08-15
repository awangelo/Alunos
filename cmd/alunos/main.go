package main

import (
	"alunos/pkg/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	// Novo mux para rotear as requisicoes, com os handlers definidos para cada metodo e rota.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /login", handlers.Login)
	mux.HandleFunc("POST /login", handlers.LoginAuth)

	// Middleware de autenticacao para as rotas protegidas.
	mux.Handle("/alunos", handlers.AuthMiddleware(http.HandlerFunc(handlers.Alunos)))
	mux.Handle("/alunos/inserir", handlers.AuthMiddleware(http.HandlerFunc(handlers.InserirAluno)))
	mux.Handle("/alunos/remover", handlers.AuthMiddleware(http.HandlerFunc(handlers.RemoverAluno)))

	// Certificado e chave.
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	log.Printf("Listening on :8080")
	if err := http.ListenAndServeTLS(":8080", certFile, keyFile, mux); err != nil {
		log.Fatalf("Nao foi possivel iniciar o servidor: %v", err)
	}
}
