package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {

	cfg := mysql.Config{
		User:                 os.Getenv("db_user"),
		Passwd:               os.Getenv("db_password"),
		Addr:                 os.Getenv("db_host"),
		DBName:               os.Getenv("db_name"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Cannot connect to database, error founded: %v", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Couldn't ping with the database, error detected: %v", err)
		return nil
	}

	return db
}
