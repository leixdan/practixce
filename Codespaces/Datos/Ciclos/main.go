package main

import "fmt"

func agregar ()  {
		
	inventary := [] string {} //Se crea el slice inventario
var item string
fmt.Println("¿Agregamos algo a las existencias?: ") //Preguntamos al usuario que desea agregar
fmt.Scan(&item)
inventary = append(inventary, item)
fmt.Println("Se ha agregado un: ",item)
fmt.Println("Inventario actual: ",inventary)

}

func main (){

	fmt.Println("Bienvenido a tu inventario")
	var respuesta string
	fmt.Println("¿Deseas agregar algún elemento a tu inventario? (y/n)")
	fmt.Scan(&respuesta)
	if respuesta == "y" {
		agregar()
	}else{
		fmt.Println("Entonces, ¿qué hacemos?")
	}

}