package persist

import (
	"database/sql"
	"embed"
	"log"
	"strings"
)

//we are probably going to want a real migrations library one day

//go:embed build.sql
var build embed.FS

func Setup(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to create a transaction :%v", err)
	}
	sb := strings.Builder{}

	b, err := build.ReadFile("build.sql")
	if err != nil {
		log.Fatalf("Failed to read in sql for database setup: %v", err)
	}
	sb.Write(b)
	_, err = tx.Exec(sb.String())
	if err != nil {
		log.Fatalf("Failed to run sql: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit: %v", err)
	}
}
