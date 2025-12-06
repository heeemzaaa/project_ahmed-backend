package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"backend/middlewares"
	m "backend/migration"
	g "backend/models"
	"backend/routes"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	os.MkdirAll("/home/hamza/Desktop/project_ahmed/backend/database", os.ModePerm)

	godotenv.Load()

	var err error
	g.DB, err = sql.Open("sqlite3", "file:/home/hamza/Desktop/project_ahmed/backend/database/database.db?_busy_timeout=2000&_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}
	m.Migrate()
}

func main() {
	mux := routes.SetRoutes(g.DB)

	log.Println("Server running on :8080")
	http.ListenAndServe("0.0.0.0:8080", middlewares.NewCorsMiddleware(mux))
}
