package main

import (
	"fmt"
)

func main() {

	//Lo primero es solicitar un numero, como int, la forma en que se solicita es declarando
	//una variable a la que no se le va a indicar un valor definido, debajo se solicita ingreso
	//por parte del usuario

	var digitouno int
	fmt.Print("Ingresa el primer valor: ")
	fmt.Scan(&digitouno)
	// En este punto se debería agregar un verificador para evitar errores por parte del usuario
	var digitodos int
	fmt.Print("Ingresa el segundo valor: ")
	fmt.Scan(&digitodos)
	//Toca sumar las dos variables
	suma := digitouno + digitodos

	//Mostrar el resultado
	fmt.Printf("La suma de %d y %d es: %d\n", digitouno, digitodos, suma)
}