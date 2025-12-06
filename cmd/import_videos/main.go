package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbPath := "/home/hamza/Desktop/project_ahmed/backend/database/database.db"
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_busy_timeout=2000&_journal_mode=WAL", dbPath))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	file, err := os.Open("./data/videos.csv")
	if err != nil {
		log.Fatal("Failed to open videos.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for i := 0; ; i++ {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading CSV:", err)
		}
		if i == 0 {
			continue // skip header
		}

		// Parse orderIndex
		orderIndex, err := strconv.Atoi(row[4])
		if err != nil {
			log.Printf("Invalid orderIndex at row %d: %v\n", i, err)
			continue
		}

		// Parse folderOrder (optional)
		var folderOrder sql.NullInt64
		if strings.TrimSpace(row[5]) == "" {
			folderOrder = sql.NullInt64{Valid: false}
		} else {
			fo, err := strconv.ParseInt(row[5], 10, 64)
			if err != nil {
				log.Printf("Invalid folderOrder at row %d: %v\n", i, err)
				folderOrder = sql.NullInt64{Valid: false}
			} else {
				folderOrder = sql.NullInt64{Int64: fo, Valid: true}
			}
		}

		_, err = db.Exec(`
			INSERT OR IGNORE INTO videos 
			(id, title, category, vdocipherVideoId, orderIndex, folderOrder)
			VALUES (?, ?, ?, ?, ?, ?)`,
			row[0], row[1], row[2], row[3], orderIndex, folderOrder,
		)
		if err != nil {
			log.Printf("Error inserting row %d: %v\n", i, err)
		}
	}

	fmt.Println("Videos import completed!")
}
