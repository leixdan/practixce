package main

import "fmt"


func main(){

	var respuesta int = 25
	var sBinario []int
	var binario int

	fmt.Println("Vamos a convertir tu número")
	
	i := respuesta
	for i > 0 {
		
		i = i/2
		}

		fmt.Println("Terminó el ciclo")
		fmt.Println(sBinario)

}