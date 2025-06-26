package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	connStr := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Error conectando: %v", err)
	}
	defer conn.Close(ctx)

	fmt.Println("Conectado con éxito a PostgreSQL")
}
