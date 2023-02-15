package db

import (
	"database/sql"
	"fmt"
	"log"

	"gpt-cms/config"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", config.Cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %s", err)
	}

	Db = db
	fmt.Println("Connected to database")
}
