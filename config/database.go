package config

import (
"database/sql"
"log"
"os"


_ "github.com/lib/pq"
"github.com/joho/godotenv"


)

func ConnectDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

db, err := sql.Open(
	"postgres",
	os.Getenv("DB_URL"),
)

if err != nil {
	log.Fatal("Failed to connect to database:", err)
}

if err := db.Ping(); err != nil {
	log.Fatal("Failed to ping database:", err)
}

log.Println("Database connected successfully")

return db


}
