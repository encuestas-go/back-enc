package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {

	cfg := mysql.Config{
		User:   "",
		Passwd: "",
		Addr:   "",
		DBName: "ENCUESTAS",
		Net:    "tcp",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Cannot connect to database, error founded: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Couldn't ping with the database, error detected: %v", err)
	}

	return db
}
