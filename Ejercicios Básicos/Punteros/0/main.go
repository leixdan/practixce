//Elaborar un programa donde se solicite al usuario el nombre a actulizr de la mascota

package main

import "fmt"

func actualizarNombre(nombre *string, nuevoValor string) { //La función recibe un puntero a string y un nuevo valor 
	*nombre = nuevoValor
}

func main() {

	var nombre string = "Firulais"
	var respuesta string
	var nuevoNombre string

	fmt.Println("Hola humano, actualmente tenemos un listado de nombres de mascotas")
	fmt.Println("Empecemos con los perros: ")
	fmt.Printf("1. %s\n", nombre)
	fmt.Println("¿Te gustaría actualizar el nombre de la mascota? (y/n)")

	fmt.Scanln(&respuesta)
	if respuesta == "y" {
		fmt.Println("Ingresa el nuevo nombre de la mascota:")
		fmt.Scanln(&nombre)
		actualizarNombre(&nombre, nuevoNombre)
		fmt.Println("El nuevo nombre de la mascota es:", nombre)
	} else {
		fmt.Println("No se actualizó el nombre de la mascota.")
	}

}
