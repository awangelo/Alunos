package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSessionToken gera um token de sessao aleatorio.
func GenerateSessionToken() string {
	// Slice of bytes para armazenar o token.
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Erro ao gerar token:", err)
		return ""
	}
	// Converte para string hex.
	return hex.EncodeToString(b)
}

// SaveSessionToken salva o token de sessao no banco de dados.
func SaveSessionToken(username string, token string) bool {
	dbPath := os.Getenv("DATABASE_PATH")
	// Abre o banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Erro ao abrir banco de dados: %v", err)
		return false
	}
	defer db.Close()

	// Insere o token de sessao no banco de dados.
	_, err = db.Exec("INSERT INTO sessions (username, token) VALUES (?, ?)", username, token)
	return err == nil
}

func ValidateLogin(username, password string) bool {
	dbPath := os.Getenv("DATABASE_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Erro ao abrir banco de dados: %v", err)
		return false
	}
	defer db.Close()

	// Query para obter o hash da senha do usuário.
	var hashedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// User nao encontrado.
			return false
		}
		log.Printf("Erro ao consultar usuário: %v", err)
		return false
	}

	// Verifica se a senha corresponde ao hash armazenado.
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func IsValidSession(token string) bool {
	dbPath := os.Getenv("DATABASE_PATH")
	// Abre o banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query para verificar se o token de sessao eh valido.
	var exists bool
	errQuery := db.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE token = ?)", token).Scan(&exists)
	if errQuery != nil {
		log.Printf("Erro ao verificar token: %v", err)
		return false
	}
	return exists
}
