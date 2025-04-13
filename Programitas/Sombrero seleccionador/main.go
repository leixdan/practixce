package main

import (
	"fmt"
)

func main() {

	//Variables
	var nombre string
	var edad int
	var resultado int //Valor apuntado
	var respuesta int

	//Bienvenida al programa
	fmt.Println("¡Bienvenido al sombrero seleccionador!")
	fmt.Println("Por favor, ingresa tu nombre: ")
	fmt.Scan(&nombre)
	fmt.Println("¿Cuántos años tienes?")
	fmt.Scan(&edad)

	//Preguntas
	fmt.Println("Por favor, responde las siguientes preguntas para determinar tu casa en Hogwarts.")

	fmt.Println("a) ¿Qué valor consideras más importante?")
	fmt.Println("1. La valentía")
	fmt.Println("2. La lealtad")
	fmt.Println("3. La inteligencia")
	fmt.Println("4. La ambición")
	fmt.Scan(&respuesta)

	fmt.Println("b) ¿Cómo prefieres resolver un conflicto?")

	fmt.Printf("El puntaje fue: %d\n", resultado)
}
