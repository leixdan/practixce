package main

import 	"fmt"

func saludar (){ //Función sin parámetros y sin retorno
	fmt.Println("Hola, este es un saludo de función simple")
}

func bienvenida (nombre string){ //Función con parámetros sin retorno
	fmt.Printf("Bienvenido, %s\n",nombre)
}

func confirmarNombre (nombre *string){ //Función con parámetros sin retorno, con puntero
	*nombre = "Mr bombastic" 
	
}

func edadUsuario (a int, b int) int{
	return a-b
}

func main(){

	saludar()

	nombre := "Lamar"
	bienvenida(nombre)

	confirmarNombre(&nombre)
	fmt.Printf("Sólo para estar seguros, ¿tu nombre es: %s?\n", nombre)

	edad := edadUsuario(2025,1995)
	fmt.Println("La edad del susodicho: ",edad)
}