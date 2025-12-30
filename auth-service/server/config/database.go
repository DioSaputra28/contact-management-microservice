package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func DbConnection() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+dbname+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
