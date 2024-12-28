package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type AccessDatabase struct {
	DB *sql.DB
}

func databaseTable(db *sql.DB) error {
	schedulerDatabaseTable := `
      CREATE TABLE IF NOT EXISTS scheduler (
        id integer PRIMARY KEY AUTOINCREMENT,
        date TEXT NOT NULL,
        title TEXT NOT NULL,
        comment TEXT NOT NULL,
        repeat VARCHAR(128) CHECK (length(repeat) <= 128)
      );
      CREATE INDEX IF NOT EXISTS index_date ON scheduler (date);
      `
	_, err := db.Exec(schedulerDatabaseTable)
	if err != nil {
		return fmt.Errorf("failed to create table and index: %w", err)
	}

	return nil
}

func OpenOrCreate(nameOfDatabase string) (*AccessDatabase, error) {
	appPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbFile := filepath.Join(appPath, nameOfDatabase)

	var install bool
	_, err = os.Stat(dbFile)
	if err != nil {

		install = true
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	if install {

		fmt.Println("Creating new database...")
		err := databaseTable(db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database and tables created.")
	} else {

		fmt.Println("Database found, proceeding...")
	}
	return &AccessDatabase{
		DB: db,
	}, nil
}
