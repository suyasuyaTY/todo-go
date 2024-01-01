package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // initialize mysql driver
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func DefaultDSN(host, port, user, password, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo", user, password, host, port, dbname)
}

func Connect() error{
	dsn := DefaultDSN(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	db = conn
	return nil
}