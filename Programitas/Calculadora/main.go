package main

import (
"calculadora/Operaciones"
"calculadora/Utils"
"fmt"
)

func main(){

	//Variables
	var resultado, op, a, b int

	//Bienvenida al programa
	fmt.Println("Bienvenido a la calculadora")
	//Alternativas al usuario
	fmt.Println("Elige la operación que deseas realizar")
	fmt.Println("1. Suma")
	fmt.Println("2. Resta")
	/*fmt.Println("3. Multiplicación")
	fmt.Println("4. División") */
	fmt.Scan(&op)

	if op != 1 || op != 2 {
		fmt.Println("Elige una opción prro")
		
	}

//Solicitud de los dos números al usuario (a y b)
fmt.Println("Ingresa el primer número: ")
fmt.Scan(&a)
fmt.Println("Ingresa el segundo número: ")
fmt.Scan(&b)

if op == 1 {
	resultado = operaciones.Sumar(a, b)

utils.ImprimirResultado(a, b, resultado)
} else if op == 2 {
	resultado = operaciones.Restar(a, b)
} else {
	fmt.Println("No hay de ese")
}
}