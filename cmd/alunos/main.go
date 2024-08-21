package main

import (
	"alunos/internal/routes"
	"log"
	"net/http"
)

func main() {
	// Novo mux para rotear as requisicoes, com os handlers definidos para cada metodo e rota.
	mux := routes.NewRouter()

	log.Printf("Listening on :80")
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("Error ao iniciar: %v", err)
	}
}
