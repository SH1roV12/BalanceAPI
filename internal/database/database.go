package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB()*sql.DB{
	godotenv.Load()
	user:=os.Getenv("DB_USER")
	pass:=os.Getenv("DB_PASS")
	host:=os.Getenv("DB_HOST")
	port:=os.Getenv("DB_PORT")
	name:=os.Getenv("DB_NAME")

	var DSN string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

	db,err:=sql.Open("mysql",DSN)
	if err!= nil{
		panic(err)
	}
	if err:=db.Ping(); err!=nil{
		panic(err)
	}
	fmt.Println("the database is open")
	return db
}