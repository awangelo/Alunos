package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Aluno struct {
	RA    int
	Email string
	M1    float64
	M2    float64
	M3    float64
}

func Alunos(w http.ResponseWriter, r *http.Request) {
	alunos := getAlunos()

	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/alunos.html"))
	err := tmpl.Execute(w, alunos)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

// getAlunos retorna uma lista de alunos do banco de dados.
func getAlunos() []Aluno {
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
