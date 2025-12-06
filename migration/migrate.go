package migration

import (
	"log"
	"os"

	g "backend/models"
)

func Migrate() {
	filePath := "./database/database.sql"

	query, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = g.DB.Exec(string(query))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migrated successfully")
}
