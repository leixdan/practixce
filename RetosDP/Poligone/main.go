package main

import "fmt"

//Funciones

func triangulo(base float64, altura float64) float64 {
	area := (base * altura) / 2
	return area
}

func cuadrado(base float64) float64 {
	area := base * base
	return area
}

func rectangulo(base float64, altura float64) float64 {
	area := base * altura
	return area
}

func main() {

	var base float64
	var altura float64
	var opcion int

	fmt.Println("¿Qué polígono quieres calcular?")
	fmt.Println("1. Triángulo")
	fmt.Println("2. Cuadrado")
	fmt.Println("3. Rectángulo")
	fmt.Print("Selecciona una opción (1-3): ")
	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		fmt.Print("Introduce la base del triángulo: ")
		fmt.Scanln(&base)
		fmt.Print("Introduce la altura del triángulo: ")
		fmt.Scanln(&altura)
		triangulo(base, altura) // Llamada a la función
		fmt.Printf("El área del triángulo es: %.2f\n", triangulo(base, altura))

	case 2:
		fmt.Print("Introduce el lado del cuadrado: ")
		fmt.Scanln(&base)
		cuadrado(base) // Llamada a la función
		fmt.Printf("El área del cuadrado es: %.2f\n", cuadrado(base))

	case 3:
		fmt.Print("Introduce la base del rectángulo: ")
		fmt.Scanln(&base)
		fmt.Print("Introduce la altura del rectángulo: ")
		fmt.Scanln(&altura)
		rectangulo(base, altura) // Llamada a la función
		fmt.Printf("El área del rectángulo es: %.2f\n", rectangulo(base, altura))
	default:
		fmt.Println("Opción no válida. Por favor, selecciona una opción entre 1 y 3.")
	}
}
