package main

import (
	"fmt"
)

func main() {

	var edad int
	fmt.Print("¿Cuál es tu edad?: ")
	fmt.Scan(&edad)

	if edad > 18 {
		var pedido string
		fmt.Print("¿De qué drogas quieres chiko?: ")
		fmt.Scan(&pedido)
		var cantidad int
		fmt.Print("Bien, bien, ¿cuántos gramos quieres?: ")
		fmt.Scan(&cantidad)
		fmt.Printf("Bien chiko, aquí tienes %d gramos de buena %s\n", cantidad, pedido)
	} else {
		fmt.Print("No tienes edad aún chiko, vuelve en unos años\n")
	}

}
