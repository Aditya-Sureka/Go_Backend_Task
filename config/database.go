package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:Aditya%407781@localhost:5432/userdb",
	)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	return conn
}