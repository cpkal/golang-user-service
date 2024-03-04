package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Run() *sql.DB {
	DB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang-user-service")

	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	return DB
}
