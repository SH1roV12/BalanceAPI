package database

import (
	"database/sql"
	"fmt"
	getconfig "users-balance/internal/getConfig"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	DSN := getconfig.GetConfig()

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("the database is open")
	return db
}
