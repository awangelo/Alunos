package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"golang.org/x/crypto/bcrypt"
)

// openDatabase abre a conexão com o banco de dados.
func openDatabase() (*sql.DB, error) {
	dbPath := os.Getenv("DATABASE_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GenerateSessionToken gera um token de sessao aleatorio.
func GenerateSessionToken() (string, error) {
	// Slice of bytes para armazenar o token.
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Erro ao gerar token:", err)
		return "", fmt.Errorf("Erro ao gerar token: %v", err)
	}
	// Converte para string hex.
	return hex.EncodeToString(b), nil
}

// SaveSessionToken salva o token de sessao no banco de dados.
func SaveSessionToken(username string, token string) bool {
	db, err := openDatabase()
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
	db, err := openDatabase()
	if err != nil {
		log.Printf("Erro ao abrir banco de dados: %v", err)
		return false
	}
	defer db.Close()

	// Query para obter o hash da senha do usuário.
	var hashedPassword string
	err = db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// User nao encontrado.
			return false
		}
		log.Printf("Erro ao consultar usuario: %v", err)
		return false
	}

	// Verifica se a senha corresponde ao hash armazenado.
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func IsValidSession(token string) bool {
	db, err := openDatabase()
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
