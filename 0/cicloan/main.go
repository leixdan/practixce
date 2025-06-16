package main

import (
	"fmt"
)

func main(){

	//Slice
	lugares := []string{"a", "b", "c", "d", "d"}
	sillas := []string{"a", "c", "b"}

	var repetidos, asignados string

	//Mapa
	analistas := map[string]string{
		
		"Gatito": "Karla",
		"Machapa": "Daniel",
		"Gato": "Aly",
		"Oso": "Iran",

	}

	//Struct
	type planta struct{
		Nombre	string
		Tamaño	float32
	}

	fmt.Println(analistas)


	for _, l := range lugares {
		encontrado := false
		for _, s := range sillas {
			if l == s {
				repetidos += l
				encontrado = true
				break // Ya no necesitamos seguir buscando
			}
		}
		if !encontrado {
			asignados += l
		}
	}

	fmt.Println("\nResultado final:")
	fmt.Println("Repetidos:", repetidos)
	fmt.Println("Asignados:", asignados)
}

