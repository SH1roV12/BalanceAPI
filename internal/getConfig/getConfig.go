package getconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//Getting config for init database
func GetConfig() string {
	//Getting .env var
	godotenv.Load()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	//Building DSN
	var DSN string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	return DSN
}
