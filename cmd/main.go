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
	os.MkdirAll("./database", os.ModePerm)

	godotenv.Load()

	var err error
	g.DB, err = sql.Open("sqlite3", "file:./database/database.db?_busy_timeout=2000&_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}
	m.Migrate()
}

func main() {
	mux := routes.SetRoutes(g.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on :8080")
	http.ListenAndServe("0.0.0.0:"+port, middlewares.NewCorsMiddleware(mux))
}
