package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

func InserirAlunoForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/inserir.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

func InserirAluno(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	ra := r.FormValue("ra")
	m1, m2, m3 := r.FormValue("m1"), r.FormValue("m2"), r.FormValue("m3")

	dbPath := os.Getenv("DATABASE_PATH")
	// Abre o banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO alunos (ra, email, nota1, nota2, nota3) VALUES (?, ?, ?, ?, ?)", ra, email, m1, m2, m3)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/alunos", http.StatusSeeOther)
}

func RemoverAluno(w http.ResponseWriter, r *http.Request) {
	// Como o metodo vai DELETE, nao utilizara parametros na URL
	// entao `ra := r.URL.Query().Get("ra")` nao funcionara.
	parts := strings.Split(r.URL.Path, "/")
	ra := parts[len(parts)-1]

	dbPath := os.Getenv("DATABASE_PATH")
	// Abre o banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query para remover o aluno.
	_, err = db.Exec("DELETE FROM alunos WHERE ra = ?", ra)
	if err != nil {
		log.Fatal(err)
	}

	// Recarrega a pagina de alunos.
	http.Redirect(w, r, "/alunos", http.StatusSeeOther)
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
