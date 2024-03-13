package db

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
    Conn *sql.DB
}

func NewMySQLDB() (*MySQLDB, error) {
    // Fill in your MySQL connection details
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang-user-service")
    if err != nil {
        return nil, err
    }

    // Check if the connection is actually working
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to MySQL database")

    return &MySQLDB{Conn: db}, nil
}
