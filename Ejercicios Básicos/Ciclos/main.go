package main

import "fmt"

func main (){



	data := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("El contenido original es: ",data)

fmt.Println("Ciclo simple")
	for i:=0; i <= 4; i++{
		if i == 2{	//Aquí nos saltamos el 2 antes de la indicación de sacarlo por la consola
			continue
		}

		fmt.Println(i)
	}
fmt.Println("____________________________")

fmt.Println("Ciclo while")//Mientras uno sea menor o igual a 5, el ciclo continuará
	uno := 0
	for uno <= 5{
		fmt.Println(uno)
		uno ++
	}
fmt.Println("____________________________")

fmt.Println("Ciclo con range") //Iterará el número de veces igual al range del slice

for i, val := range data{

	if i == 3 {	//Esta condicional comparará el índice y saldrá del ciclo
		break
	}

	fmt.Printf("Index: %d, Value: %s\n",i ,val)
}
fmt.Println("____________________________")



}