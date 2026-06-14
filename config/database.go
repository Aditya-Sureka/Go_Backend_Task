package config

import (
"database/sql"
"log"


_ "github.com/lib/pq"


)

func ConnectDB() *sql.DB {


db, err := sql.Open(
	"postgres",
	"postgres://postgres:Aditya%407781@localhost:5432/userdb?sslmode=disable",
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
