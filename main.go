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
	ID       int
	Ev       sql.NullInt64
	Felev    sql.NullInt64
	Sorszam  sql.NullInt64
	Kerdes   string
	Megoldas string
	Generalt bool
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
		err := db.QueryRow("SELECT * FROM erettsegizteto.feladat WHERE id = $1", rand.Intn(5)+1).Scan(&row.ID, &row.Ev, &row.Felev, &row.Sorszam, &row.Kerdes, &row.Megoldas, &row.Generalt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute HTML template
		err = tmpl.Execute(w, map[string]interface{}{
			"Kerdes":   row.Kerdes,
			"Ev":       nullToString(row.Ev),
			"Felev":    nullToString(row.Felev),
			"Sorszam":  nullToString(row.Sorszam),
			"Generalt": row.Generalt,
			"Method":   r.Method,
			"Answer":   r.FormValue("answer"),
			"Megoldas": row.Megoldas,
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
