package main

import (
	operaciones "calculadora/Operaciones"
	utils "calculadora/Utils"
	"fmt"
)

func main() {

	//Variables
	var op int
	var a, b, resultado float64

	//Bienvenida al programa
	fmt.Println("Bienvenido a la calculadora")
	//Alternativas al usuario
	fmt.Println("Elige la operación que deseas realizar")
	fmt.Println("1. Suma")
	fmt.Println("2. Resta")
	fmt.Println("3. Multiplicación")
	fmt.Println("4. División")
	fmt.Scan(&op)

	switch op {
	case 1:
		fmt.Println("Ingresa el primer número: ")
		fmt.Scan(&a)
		fmt.Println("Ingresa el segundo número: ")
		fmt.Scan(&b)
	case 2:
		fmt.Println("Ingresa el primer número: ")
		fmt.Scan(&a)
		fmt.Println("Ingresa el segundo número: ")
		fmt.Scan(&b)
	case 3:
		fmt.Println("Ingresa el primer número: ")
		fmt.Scan(&a)
		fmt.Println("Ingresa el segundo número: ")
		fmt.Scan(&b)
	case 4:
		fmt.Println("Recuerda que no se puede dividir entre 0")
		fmt.Println("Ingresa el primer número: ")
		fmt.Scan(&a)
		fmt.Println("Ingresa el segundo número: ")
		fmt.Scan(&b)
	default:
		fmt.Println("Elige una opción correcta prro")
		utils.TerminarPrograma()
	}

	if op == 1 {
		resultado = operaciones.Sumar(a, b)
		utils.ImprimirResultado(a, b, resultado)
	} else if op == 2 {
		resultado = operaciones.Restar(a, b)
		utils.ImprimirResultado(a, b, resultado)
	} else if op == 3 {
		resultado = operaciones.Multiplicar(a, b)
		utils.ImprimirResultado(a, b, resultado)
	} else if op == 4 {
		resultado = operaciones.Dividir(a, b)
		utils.ImprimirResultado(a, b, resultado)
	} else {
		fmt.Println("No hay de ese")
	}
}
