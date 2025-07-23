package main

import (
	"fmt"
)

func main() {

	fmt.Println("Este programa te ayudará a decidir si llevar o no paraguas hoy")

	// Ingreso de datos del usuario
	fmt.Println("¿Está nublado afuera? (y/n): ")
	var r1 string
	fmt.Scan(&r1)

	if r1 == "y" {
		fmt.Println("¿Cuántas horas estarás fuera?")
		var r2 int
		fmt.Scan(&r2)

		if r1 == "y" && r2 <= 3 {
			fmt.Println("No hace falta")
		} else if r1 == "y" && r2 <= 5 {
			fmt.Println("Si, deberías llevarlo")
		} else {
			fmt.Println("Pide un uber o mejor no salgas")
		}
	} else if r1 == "n" {
		fmt.Println("No tendrías que estar usando este programa...")
	} else {
		fmt.Println("Ingresa una respuesta válida: y / n")
	}

}
