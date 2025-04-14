//Ejercicios mamalones de punteros porque estoy bien wey y se me olvidan las funciones tambien

package main

import "fmt"

func duplicar(n int) int { //Ej 1
	return n * 2
}

func triplicarPuntero(n *int) { //Ej 2
	*n = *n * 3
}

func main() {

	//Variables

	var dup int

	//Ejercicio 1
	fmt.Println("Ingresa el número que deseas duplicar: ")
	fmt.Scan(&dup)
	resultado := duplicar(dup)
	fmt.Println("El resultado es:", resultado)

	//Ejercicio 2
	fmt.Println("Número triplicado por puntero: ")
	triplicarPuntero(&dup)
	fmt.Println("El resultado es:", dup)
}
