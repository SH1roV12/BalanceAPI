package database

import (
	"database/sql"
	"log"
	getconfig "users-balance/internal/getConfig"

	_ "github.com/go-sql-driver/mysql"
)

//Initialization of database
func InitDB() *sql.DB {
	//Getting config for open database
	DSN := getconfig.GetConfig()

	//Opening database
	db, err := sql.Open("mysql", DSN)

	
	if err != nil {
		log.Fatalln("Cannot open database")
	}


	if err := db.Ping(); err != nil {
		log.Fatalln("Cannot ping database")
	}


	log.Println("The database is open")
	return db
}
