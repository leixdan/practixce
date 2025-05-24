package main

import "fmt"

func main (){

	var revisar int

	fmt.Println("Introduce el número a revisar chiko: ")
	fmt.Scan(&revisar)

	if revisar == 0 { //La solución más rápida era evitar el 0, aunque con manejo de errores me evito el if anidado
			println("El cero no cuenta")
		} else {
			if revisar % 2 == 0 {
		println("Este número es par")
		} else {
		println("Este número es impar")
		}
	}
}