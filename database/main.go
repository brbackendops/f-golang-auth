package models

import (
	"fmt"
	"log"
	"os"

	"falcon/database/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func Connect() *sqlx.DB {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbName, sslMode,
	))

	if err != nil {
		log.Fatalln(err.Error())
	}

	tx := db.MustBegin()
	tx.MustExec(models.Schema)
	tx.Commit()

	defer tx.Rollback()

	DB = db
	return db
}

func ConnectTestDb() *sqlx.DB {

	db, err := sqlx.Connect("sqlite3", ":memory:?cache=shared")

	if err != nil {
		fmt.Println("from db", err)
		log.Fatalln(err.Error())
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error from test db", err)
	}

	schema := `
	
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at	DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`

	tx := db.MustBegin()
	tx.MustExec(schema)
	tx.Commit()

	DB = db
	return db
}
