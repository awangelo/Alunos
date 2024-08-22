package models

import (
	"database/sql"
	"log"
	"os"
)

type Aluno struct {
	RA    int
	Email string
	M1    float64
	M2    float64
	M3    float64
}

// GetAlunos retorna uma lista de todos os alunos registrados no banco de dados.
func GetAlunos() []Aluno {
	dbPath := os.Getenv("DATABASE_PATH")
	// Abre o banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query para obter os alunos.
	rows, err := db.Query("SELECT ra, email, nota1, nota2, nota3 FROM alunos")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var alunos []Aluno
	for rows.Next() {
		var aluno Aluno
		err := rows.Scan(&aluno.RA, &aluno.Email, &aluno.M1, &aluno.M2, &aluno.M3)
		if err != nil {
			log.Fatal(err)
		}
		alunos = append(alunos, aluno)
	}

	return alunos
}

// GetAluno retorna um aluno do banco de dados utilizando o RA como chave.
func GetAluno(ra string) (Aluno, error) {
	dbPath := os.Getenv("DATABASE_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return Aluno{}, err
	}
	defer db.Close()

	var aluno Aluno
	err = db.QueryRow("SELECT ra, email, nota1, nota2, nota3 FROM alunos WHERE ra = ?", ra).Scan(&aluno.RA, &aluno.Email, &aluno.M1, &aluno.M2, &aluno.M3)
	if err != nil {
		return Aluno{}, err
	}

	return aluno, nil
}

// GetNumAlunos retorna o numero de alunos registrados.
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

// InsertAluno insere um aluno e todos os seus dados no banco de dados.
func InsertAluno(ra, email, m1, m2, m3 string) error {
	dbPath := os.Getenv("DATABASE_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO alunos (ra, email, nota1, nota2, nota3) VALUES (?, ?, ?, ?, ?)", ra, email, m1, m2, m3)
	return err
}

// UpdateAluno atualiza os dados de um aluno no banco de dados.
func UpdateAluno(ra, email, m1, m2, m3 string) error {
	dbPath := os.Getenv("DATABASE_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE alunos SET email = ?, nota1 = ?, nota2 = ?, nota3 = ? WHERE ra = ?", email, m1, m2, m3, ra)
	return err
}

// DeleteAluno remove um aluno do banco de dados utilizando o RA como chave.
func DeleteAluno(ra string) error {
	dbPath := os.Getenv("DATABASE_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM alunos WHERE ra = ?", ra)
	return err
}
