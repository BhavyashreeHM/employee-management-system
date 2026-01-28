package database

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"


	"github.com/joho/godotenv"
)

func ConnectSQL() (sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return sql.DB{}, err
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return sql.DB{}, err
	}

	if err := db.Ping(); err != nil {
		return sql.DB{}, err
	}

	return *db, nil

}
