package main

import (
	"http/handlers"
	"log"
	"net/http"
)

// 192.168.15.123
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/registro", handlers.Registro)
	http.HandleFunc("/alunos", handlers.Alunos)
	http.HandleFunc("/login", handlers.Login)

	log.Printf("Starting server on port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
