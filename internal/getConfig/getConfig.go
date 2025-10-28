package getconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() string {
	godotenv.Load()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	var DSN string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	return DSN
}
