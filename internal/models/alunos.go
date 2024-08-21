package models

import (
	"database/sql"
	"log"
	"os"
)

func GetNumAlunos() int {
	dbPath := os.Getenv("DATABASE_PATH")
	// Abre o banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query para obter o numero de alunos.
	var numAlunos int
	// Ira retornar apenas uma linha com o numero de alunos.
	err = db.QueryRow("SELECT COUNT(*) FROM alunos").Scan(&numAlunos)
	if err != nil {
		log.Fatal(err)
	}

	return numAlunos
}
