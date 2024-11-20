package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"product-management/configs"
)

var db *sql.DB

func ConnectDB(cfg *configs.Config) {
	var err error
	db, err = sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Connected to PostgreSQL database!")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
