package main

import "fmt"


func main(){

	var respuesta int
	var sBinario []int

	fmt.Println("Vamos a convertir tu número")
	fmt.Println("Ingresa el número que deseas convertir a binario: ")
	fmt.Scan(&respuesta)
	
	i := respuesta
	for i > 0 {
		if i % 2 == 0 {
			sBinario = append(sBinario, 0)
		} else {
			sBinario = append(sBinario, 1)
		}
		i = i/2

		//sBinario = append(sBinario, i)
		}
		var invertido string
		for i := len(sBinario) - 1; i >= 0; i--{
//Se hace uso de Sprintf para convertir los números
//a strings ya que si se usa string suma los 1 y arroja
//espacios vacíos, no es lo buscado
			invertido += fmt.Sprintf("%d",sBinario[i])
		}
		fmt.Println("Terminó el ciclo")
		fmt.Printf("Tu número en binario es: %s\n",invertido)

}