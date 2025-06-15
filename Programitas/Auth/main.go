package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Aquí pones tu conexión a la base de datos
	dsn := "postgres://authuser:deckna7749@localhost:5432/login_db"

	// Creamos un contexto para la conexión
	ctx := context.Background()

	// Conectamos a la base de datos
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer conn.Close(ctx)

	fmt.Println("¡Conexión exitosa!")
}
