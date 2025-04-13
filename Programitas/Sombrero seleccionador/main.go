package main

import (
	"fmt"
)

func main() {

	//Variables
	var nombre string
	var edad, respuestas int

	//Bienvenida al programa
	fmt.Println("¡Bienvenido al sombrero seleccionador!")
	fmt.Println("Por favor, ingresa tu nombre: ")
	fmt.Scan(&nombre)
	fmt.Println("¿Cuántos años tienes?")
	fmt.Scan(&edad)

	//Preguntas
	fmt.Println("Por favor, responde las siguientes preguntas para determinar tu casa en Hogwarts.")

	fmt.Println("1. ¿Qué valor consideras más importante?")

}
