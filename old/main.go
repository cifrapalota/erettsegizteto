package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	_ "github.com/lib/pq"
)

type Feladat struct {
	ID        int
	Question  string
	Answer    string
	Generated bool
	Year      sql.NullInt64
	Semester  sql.NullInt64
	Number    sql.NullInt64
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	// Connect to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize rand
	rand.Seed(time.Now().Unix())

	// Parse HTML template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Create HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get a random row from the database
		row := Feladat{}
		err := db.QueryRow("SELECT * FROM db_erettsegizteto.t_question WHERE id = $1",
			rand.Intn(3)+1).Scan(&row.ID, &row.Question, &row.Answer, &row.Generated, &row.Year, &row.Semester, &row.Number)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute HTML template
		err = tmpl.Execute(w, map[string]interface{}{
			"Kerdes":   row.Question,
			"Ev":       nullToString(row.Year),
			"Felev":    nullToString(row.Semester),
			"Sorszam":  nullToString(row.Number),
			"Generalt": row.Generated,
			"Method":   r.Method,
			"Answer":   r.FormValue("answer"),
			"Megoldas": row.Answer,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}

// nullToString converts a sql.NullInt64 to a string, returning an empty string if the value is null.
func nullToString(n sql.NullInt64) string {
	if n.Valid {
		return strconv.FormatInt(n.Int64, 10)
	}
	return ""
}
