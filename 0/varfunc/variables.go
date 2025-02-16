package main

import "fmt"

//sumar suma dos números
func sumar(a, b float64) float64 {
	return a + b
}
//restar resta dos números
func restar(a, b float64) float64 {
	return a - b
}
//multiplicar multiplica dos números
func multiplicar(a, b float64) float64 {
	return a * b
}
//dividir divide dos números
func dividir(a, b float64) float64 {
	return a / b
}
 func main() {

	for {
		
	fmt.Println("Bienvenido a la calculadora Gaskull 2.0")
	fmt.Println("¿Qué operación deseas realizar?")
	fmt.Println("1. Sumar")
	fmt.Println("2. Restar")
	fmt.Println("3. Multiplicar")
	fmt.Println("4. Dividir")
	fmt.Println("5. Salir")

	var operacion int
	_, err := fmt.Scanln(&operacion)
	if err != nil {
		fmt.Println("Error al leer la opción elegida, ingresa un número")
		fmt.Scanln() //Limpiar el buffer de entrada
		continue
	}
	if operacion == 5 {
		fmt.Println("Gracias por usar la calculadora Gaskull 2.0")
		break
	}
	fmt.Println("Ingresa el primer número")
	var a float64
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Error al leer el número")
		return //Termina la ejecución del programa
	}
	fmt.Println("Ingresa el segundo número")
	var b float64
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Error al leer el número")
		return //Termina la ejecución del programa
	}

	switch operacion {
	case 1:
		fmt.Println("El resultado de la suma es: ", sumar(a, b))
	case 2:
		fmt.Println("El resultado de la resta es: ", restar(a, b))
	case 3:
		fmt.Println("El resultado de la multiplicación es: ", multiplicar(a, b))
	case 4:
		fmt.Println("El resultado de la división es: ", dividir(a, b))
	default:
		fmt.Println("Operación no válida")
	}
	
	}
 }